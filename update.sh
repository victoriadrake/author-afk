GOOS=linux go build && \
zip authorAFK.zip authorAFK && \
rm authorAFK && \
aws lambda update-function-code --function-name authorAFK --zip-file fileb://authorAFK.zip && \
rm authorAFK.zip