package utils

import (
	"context"
	"crypto/rand"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type FunctionHandlerProvider interface {
	Write(ctx context.Context, outputPath string, handlerFileName string, functionConfiguration FunctionConfiguration) (error)
}

type AwsFunctionHandlerProvider struct {

}

func (p *AwsFunctionHandlerProvider) Write( outputPath string, handlerFileName string, functionConfiguration FunctionConfiguration) (error) {

	streamifyOverrideFileContent := `global.awslambda = {
        streamifyResponse: function (handler) {
                return async (event, context) => {
                        await handler(event, event.responseStream, context);
                }
        },
};`

	randomFileId := make([]byte, 8)
	_, err := rand.Read(randomFileId)
	if err != nil {
		return err
	}

	handlerContent := fmt.Sprintf(`import './setupLambdaGlobals_%s.mjs';
	import { %s as genezioDeploy } from "./%s";
	
	function formatTimestamp(timestamp) {
	  const date = new Date(timestamp);
	
	  const day = String(date.getUTCDate()).padStart(2, "0");
	  const monthNames = [
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec"
	  ];
	  const month = monthNames[date.getUTCMonth()];
	  const year = date.getUTCFullYear();
	
	  const hours = String(date.getUTCHours()).padStart(2, "0");
	  const minutes = String(date.getUTCMinutes()).padStart(2, "0");
	  const seconds = String(date.getUTCSeconds()).padStart(2, "0");
	
	  const formattedDate = ` +
				"`${day}/${month}/${year}:${hours}:${minutes}:${seconds} +0000`" +
				`;
	  return formattedDate;
	}
	
	const handler = async function(event) {
	  const http2CompliantHeaders = {};
	  for (const header in event.headers) {
		http2CompliantHeaders[header.toLowerCase()] = event.headers[header];
	  }
	
	  const req = {
		version: "2.0",
		routeKey: "$default",
		rawPath: event.url.pathname,
		rawQueryString: event.url.search,
		headers: http2CompliantHeaders,
		queryStringParameters: Object.fromEntries(event.url.searchParams),
		requestContext: {
		  accountId: "anonymous",
		  apiId: event.headers.Host.split(".")[0],
		  domainName: event.headers.Host,
		  domainPrefix: event.headers.Host.split(".")[0],
		  http: {
			method: event.http.method,
			path: event.http.path,
			protocol: event.http.protocol,
			sourceIp: event.http.sourceIp,
			userAgent: event.http.userAgent
		  },
		  requestId: "undefined",
		  routeKey: "$default",
		  stage: "$default",
		  time: formatTimestamp(event.requestTimestampMs),
		  timeEpoch: event.requestTimestampMs
		},
		body: event.isBase64Encoded
		  ? Buffer.from(event.body, "base64")
		  : event.body.toString(),
		isBase64Encoded: event.isBase64Encoded,
		responseStream: event.responseStream,
	  };
	
	  const result = await genezioDeploy(req).catch(error => {
		console.error(error);
		return {
		  statusCode: 500,
		  body: "Internal server error"
		};
	  });
	
	  return result;
	};
	
	export { handler };`, randomFileId, functionConfiguration.Handler, functionConfiguration.Entry)

	err = WriteToFile(outputPath,handlerFileName,handlerContent,false)
	if err != nil {
		return err
	}

	err = WriteToFile(outputPath,fmt.Sprintf("setupLambdaGlobals_%s.mjs",randomFileId),streamifyOverrideFileContent,false)
	if err != nil {
		return err
	}

	return nil

}

func FunctionToCloudInput(functionElement FunctionConfiguration, backendPath string) (GenezioCloudInput, error) {
	handlerProvider := AwsFunctionHandlerProvider{}

	tmpFolderPath, err := CreateTemporaryFolder(nil, nil);
	if err != nil {
		return GenezioCloudInput{}, err
	}

	archivePath := filepath.Join(tmpFolderPath, "genezioDeploy.zip")

	err = CopyFileOrFolder(filepath.Join(backendPath, functionElement.Path), tmpFolderPath)
	if err != nil {
		return GenezioCloudInput{}, err
	}


	unzippedBundleSize,err :=  GetBundleFolderSizeLimit(tmpFolderPath)
	if err != nil {
		return GenezioCloudInput{}, err
	}

	entryFileName := "index.mjs"

	for _, err := os.Stat(filepath.Join(tmpFolderPath, entryFileName)); !os.IsNotExist(err);{
		randomName := make([]byte, 6)
		_, err := rand.Read(randomName)
		
		if err != nil {
			return GenezioCloudInput{},err
		}
		tmpName := fmt.Sprintf("%x", randomName)
		entryFileName = fmt.Sprintf("index-%s.mjs",tmpName )

	}

	err = handlerProvider.Write(tmpFolderPath, entryFileName, functionElement)
	if err != nil {
		return GenezioCloudInput{}, err
	}

	exclussionList := []string{".git",".github"}
	err = ZipDirectory(tmpFolderPath, archivePath,exclussionList)
	if err != nil {
		return GenezioCloudInput{}, err
	}

	return GenezioCloudInput{
		Type: "function",
		Name: functionElement.Name,
		ArchivePath: archivePath,
		EntryFile: entryFileName,
		UnzippedBundleSize: unzippedBundleSize,
	}, nil
}

func GetBundleFolderSizeLimit(directoryPath string) (int64, error) {
	var totalSize int64

	err := filepath.Walk(directoryPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("error walking through directory: %v", err)
	}

	fmt.Printf("Total size of the bundle: %d bytes\n", totalSize)
	return totalSize, nil
}



