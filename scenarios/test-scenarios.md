# Scenarios (combination of pod number, sizes and machine instances) to test

1. Basic case, if one instance with sufficient capacity is available
 pods: (1 CPU, 1GB) x 3, machine instance: m5.2xlarge (8CPU, 32GB)
2. If two instance types are available, and cost of using one large instance instead of multiple small ones, is lesser or nearly equal
pods: (0.5CPU, 1 GB RAM) x 4, machine instances: m5.large(2CPU,8GB,$31.556),  m5.xlarge(4CPU,16GB,$63.139) <-- in this case, cost of larger instance is marginally greater than 2 smaller instances
3. If CPU wastage is same in two instance types, but memory wastage and cost differs
pods:(0.5CPU, 1 GB RAM), machine instances: m5.large(2CPU,8GB,$31.556), a1.large(2CPU,4GB,$15.972)
4. If Memory wastage is the same in two instances, but CPU wastage and cost differs
pods: (1CPU, 1 GB RAM)x 2, machine instances: c5.2xlarge(8CPU,16GB,$106.5), m5.xlarge(4CPU,16GB,$63.139)
5. If memory and CPU are the same, but cost differs
pods: (1CPU, 1 GB RAM)x 2, machine instances: m5.xlarge(4CPU,16GB,$63.139), m4.xlarge(4CPU,16GB,$76.861)
6. When bin-packing order matters:
pods: (1CPU, 0.5GB RAM) x2, (2CPU, 0.5GB RAM), machine instances: c3.large(2CPU,3.8GB,$41.194), c3.xlarge(4CPU,7.5GB,$81.056)


## Notes

1. Need to consider additional capacity for daemon sets running on the nodes
2. Another thing to consider: because daemon set runs per node, fewer nodes will result in less resource use by daemon sets, in absolute terms