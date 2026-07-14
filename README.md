# cpu-usage
**Adaptive CPU Resource Management in Distributed Systems**

### Paper Information
- **Author(s):** SaiKrishna Mylavarapu
- **Published In:** International Journal of Intelligent Systems and Applications in Engineering (IJISAE)
- **Publication Date:** June, 2022
- **ISSN:** E-ISSN: 2147-6799
- **DOI:**
- **Impact Factor:** 

### Abstract
Efficient CPU resource management is essential for maintaining high performance in distributed systems as workloads and cluster sizes increase. Static workload allocation results in uneven CPU utilization, while conventional load balancing techniques cannot effectively adapt to dynamic workload variations. The study analyzes CPU usage across clusters with 3, 5, 7, 9, and 11 nodes to evaluate the limitations of existing allocation strategies. It demonstrates how workload imbalance and coordination overhead reduce resource utilization, system efficiency, and scalability. The findings highlight the need for adaptive workload allocation mechanisms based on real-time CPU usage to achieve balanced resource utilization and improved overall performance.

### Core Technical Contributions

- **Adaptive CPU Resource Allocation Framework:** Designed a dynamic workload allocation framework that distributes tasks based on real time CPU utilization, improving resource efficiency and balanced node utilization.

- **Intelligent Load Balancing Mechanism:** Implemented a load balancing strategy that selects the least loaded active server, reducing workload imbalance and improving CPU utilization across distributed nodes.

- **Adaptive Auto Scaling Strategy:** Developed an auto scaling mechanism that provisions additional worker nodes when CPU load exceeds predefined thresholds, ensuring sustained performance under increasing workloads.

- **Distributed System Simulation Model:** Implemented a Go based distributed system simulator integrating load balancing, server health monitoring, and adaptive scaling to evaluate CPU resource management strategies.

- **Scalability Analysis Across Cluster Sizes:** Evaluated CPU utilization across clusters of 3, 5, 7, 9, and 11 nodes, demonstrating the scalability benefits of adaptive resource allocation over deterministic approaches. 







- **Improved Transaction Processing Efficiency:**  
Processing multiple transactions together reduces communication rounds and synchronization delays, enabling faster completion of distributed transactions and better utilization of system resources.

- **Better Scalability for Distributed Systems:**  
Commit latency grows gradually with cluster expansion because fewer coordination rounds are required, enabling distributed systems to maintain stable transaction processing performance.

- **Applicability to Distributed Platforms:**  
The approach benefits distributed databases, cloud transaction systems, financial platforms, and microservice architectures that require efficient transaction completion and scalable commit processing.
 
### Experimental Results (Summary)

  | Nodes | Immediate Commit | Group Commit | Improvment (%) |
  |-------|------------------| -------------| ---------------|
  | 3     |  160             | 95           | 40.63          |
  | 5     |  185             | 110          | 40.54          |
  | 7     |  210             | 125          | 40.48          |
  | 9     |  235             | 140          | 40.43          |
  | 11    |  260             | 155          | 40.38          |

### Citation
Transaction Batching for Low Latency Commit Processing in Distributed Systems
* SaiKrishna Mylavarapu
* International Journal on Science and Technology (IJSAT) 
* ISSN E-ISSN: 2229-7677
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org \
**Author Contact** \
**LinkedIn**: linkedin.com/in/saikrishna-mylavarapu-35479114 | **Email**: krishnamysap@gmail.com






