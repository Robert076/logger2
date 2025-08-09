resource "aws_iam_role" "robert_eks_control_plane" {
  name = "RobertEKS-ControlPlane"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "eks.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "eks_cluster_policy_robert" {
  role       = aws_iam_role.robert_eks_control_plane.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
}

resource "aws_iam_role_policy_attachment" "eks_vpc_resource_controller_robert" {
  role       = aws_iam_role.robert_eks_control_plane.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController"
}

resource "aws_iam_role" "robert_eks_worker_node" {
  name = "RobertEKS-WorkerNode"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "AmazonEKSWorkerNodePolicy_robert" {
  role       = aws_iam_role.robert_eks_worker_node.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
}

resource "aws_iam_role_policy_attachment" "AmazonEC2ContainerRegistryReadOnly_robert" {
  role       = aws_iam_role.robert_eks_worker_node.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
}

resource "aws_iam_role_policy_attachment" "AmazonEKS_CNI_Policy_robert" {
  role       = aws_iam_role.robert_eks_worker_node.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
}

resource "aws_eks_cluster" "roberts_first_cluster" {
  name     = "RobertsFirstCluster"
  role_arn = aws_iam_role.robert_eks_control_plane.arn

  vpc_config {
    subnet_ids = [
      aws_subnet.subnet_logger2_api.id,
      aws_subnet.subnet_logger2_logger.id
    ]
  }

  depends_on = [
    aws_iam_role_policy_attachment.eks_cluster_policy_robert,
    aws_iam_role_policy_attachment.eks_vpc_resource_controller_robert
  ]
}

resource "aws_eks_node_group" "roberts_nodes" {
  cluster_name    = aws_eks_cluster.roberts_first_cluster.name
  node_group_name = "roberts-node-group"
  node_role_arn   = aws_iam_role.robert_eks_worker_node.arn
  subnet_ids = [
    aws_subnet.subnet_logger2_api.id,
    aws_subnet.subnet_logger2_logger.id
  ]
  scaling_config {
    desired_size = 2
    max_size     = 3
    min_size     = 1
  }

  ami_type       = "BOTTLEROCKET_x86_64"
  instance_types = ["t3.medium"]

  depends_on = [
    aws_iam_role_policy_attachment.AmazonEC2ContainerRegistryReadOnly_robert,
    aws_iam_role_policy_attachment.AmazonEKS_CNI_Policy_robert,
    aws_iam_role_policy_attachment.AmazonEKSWorkerNodePolicy_robert
  ]
}
