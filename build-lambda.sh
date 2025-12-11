#!/bin/bash

echo "🔨 Building Lambda function for HyperShop..."

# Build for Linux AMD64 (Lambda environment)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap cmd/lambda/main.go

if [ $? -eq 0 ]; then
    echo "✅ Build successful: bootstrap"
    
    # Create ZIP file
    zip -j function.zip bootstrap
    
    echo "✅ Deployment package created: function.zip"
    echo "📦 Size: $(du -h function.zip | cut -f1)"
    
    # Clean up
    rm bootstrap
    
    echo ""
    echo "🚀 Ready to deploy!"
    echo "📝 Next steps:"
    echo "   1. Go to AWS Lambda Console"
    echo "   2. Create/Update function"
    echo "   3. Upload function.zip"
else
    echo "❌ Build failed"
    exit 1
fi
