# tax-calculator

Starter project for deploying a Vite + React frontend to Google Cloud Run with Terraform.

## Project structure

```text
.
├── Dockerfile
├── nginx.conf
├── package.json
├── index.html
├── src/
│   ├── App.css
│   ├── App.jsx
│   └── main.jsx
└── terraform/
    ├── main.tf
    ├── outputs.tf
    ├── providers.tf
    ├── terraform.tfvars
    ├── variables.tf
    └── versions.tf
```

## Run React app locally

```bash
npm install
npm run dev
```

## Build and run with Docker

```bash
docker build -t tax-calculator:local .
docker run --rm -p 8080:8080 tax-calculator:local
```

Then open `http://localhost:8080`.

## Terraform usage

```bash
cd terraform
cp terraform.tfvars
terraform init
terraform plan
terraform apply
```

## gcloud usage

```bash
gcloud init
gcloud auth login
gcloud auth application-default login
```
## Create Service Account Key

```bash
export GOOGLE_CLOUD_PROJECT=financial-tools-502613
gcloud iam service-accounts keys create ~/terraform-key.json \
  --iam-account=terraform-service-account@$GOOGLE_CLOUD_PROJECT.iam.gserviceaccount.com
export GOOGLE_APPLICATION_CREDENTIALS="/home/node/terraform-key.json"
```

## Manual Cloud Run deployment flow

1. Authenticate Docker with Artifact Registry (one-time setup per environment):

```bash
   gcloud auth configure-docker australia-southeast1-docker.pkg.dev
```

2. Build and push the image:

```bash
   npm run gcp:deploy
```

   This runs `gcp:build` and `gcp:push` in sequence. Equivalent raw commands:

```bash
   docker build -t australia-southeast1-docker.pkg.dev/financial-tools-502613/financial-tools/tax-calculator:latest .
   docker push australia-southeast1-docker.pkg.dev/financial-tools-502613/financial-tools/tax-calculator:latest
```
   ```

3. Ensure `terraform/terraform.tfvars` has:
   - `project_id = "financial-tools-502613"`
   - `region = "australia-southeast1"`
   - `container_image` matching the pushed image URI.

4. Apply Terraform:

   ```bash
   cd terraform
   terraform apply
   ```

5. Retrieve the deployed URL:

   ```bash
   terraform output cloud_run_url
   ```
