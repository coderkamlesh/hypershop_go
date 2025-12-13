Write-Host " Building Lambda function for HyperShop (ARM64)..." -ForegroundColor Cyan

# Set environment variables for cross-compilation
$env:GOOS = "linux"
$env:GOARCH = "arm64"   # <--- Ye Change kiya hai (Important for AWS Graviton)
$env:CGO_ENABLED = "0"

# Build
Write-Host "Building Go binary..." -ForegroundColor Yellow
# Note: Ensure path to main.go is correct. Using ./cmd/lambda/main.go based on your structure
go build -tags lambda.norpc -o bootstrap cmd/lambda/main.go

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build successful: bootstrap" -ForegroundColor Green
    
    # Create ZIP file
    Write-Host "Creating deployment package..." -ForegroundColor Yellow
    
    if (Test-Path function.zip) {
        Remove-Item function.zip
    }
    
    # Compress-Archive command
    Compress-Archive -Path bootstrap -DestinationPath function.zip -Force
    
    # Get file size
    $size = (Get-Item function.zip).Length / 1MB
    Write-Host " Deployment package created: function.zip" -ForegroundColor Green
    Write-Host " Size: $([math]::Round($size, 2)) MB" -ForegroundColor Green
    
    # Clean up (Delete the binary, keep the zip)
    Remove-Item bootstrap
    
    Write-Host ""
    Write-Host "Ready to deploy!" -ForegroundColor Green
    Write-Host "Next steps:" -ForegroundColor Cyan
    Write-Host "   1. Go to AWS Lambda Console" -ForegroundColor White
    Write-Host "   2. Upload function.zip" -ForegroundColor White
    Write-Host "   3. Test again!" -ForegroundColor White
} else {
    Write-Host " Build failed" -ForegroundColor Red
    exit 1
}