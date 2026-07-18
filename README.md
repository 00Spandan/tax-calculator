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
    ├── terraform.tfvars.example
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
cp terraform.tfvars.example terraform.tfvars
terraform init
terraform plan
terraform apply
```

## Manual Cloud Run deployment flow

1. Build and push the image to Artifact Registry:

   ```bash
   gcloud auth configure-docker australia-southeast1-docker.pkg.dev
   docker build -t australia-southeast1-docker.pkg.dev/financial-tools/tax-calculator/tax-calculator:latest .
   docker push australia-southeast1-docker.pkg.dev/financial-tools/tax-calculator/tax-calculator:latest
   ```

2. Ensure `terraform/terraform.tfvars` has:
   - `project_id = "financial-tools"`
   - `region = "australia-southeast1"`
   - `container_image` matching the pushed image URI.

3. Apply Terraform:

   ```bash
   cd terraform
   terraform apply
   ```

4. Retrieve the deployed URL:

   ```bash
   terraform output cloud_run_url
   ```
