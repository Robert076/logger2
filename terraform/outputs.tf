output "vpc_id" {
  value = aws_vpc.roberts_vpc.id
}

output "subnet_ids" {
  value = [
    aws_subnet.subnet_logger2_api.id,
    aws_subnet.subnet_logger2_logger.id
  ]
}
