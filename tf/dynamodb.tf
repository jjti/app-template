module "dynamodb-table" {
  source  = "terraform-aws-modules/dynamodb-table/aws"
  version = "3.3.0"

  name     = "echo-table"
  hash_key = "id"
  deletion_protection_enabled = false

  attributes = [
    {
      name = "id"
      type = "S"
    },
  ]
}
