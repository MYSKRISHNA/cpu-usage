# cpu-usage
**Adaptive CPU Resource Management in Distributed Systems**

### Paper Information
- **Author(s):** SaiKrishna Mylavarapu
- **Published In:** International Journal of Intelligent Systems and Applications in Engineering (IJISAE)
- **Publication Date:** March, 2023
- **ISSN:** E-ISSN: 2147-6799
- **DOI:**
- **Impact Factor:** 

### Abstract
Distributed transaction systems often experience high commit latency because each transaction is processed independently with repeated coordination and synchronization among nodes. This work examines the impact of immediate commit processing on latency as cluster size increases. A batching based commit approach is introduced where multiple transactions are processed collectively to reduce repeated coordination overhead. Experimental analysis across different cluster sizes shows that grouped commit processing significantly lowers commit latency and improves scalability in distributed environments.

### Core Technical Contributions
- **Batch Based Commit Processing Approach:**  
Introduced a commit method that groups multiple transactions into a single coordination cycle, reducing repeated commit operations and lowering synchronization overhead in distributed transaction systems.

- **Reduced Coordination Overhead:**  
Designed a commit model that minimizes repeated communication and synchronization between coordinator and participant nodes, improving efficiency during the commit phase of transactions.

- **Distributed Transaction Simulation Model:** 
Implemented a distributed transaction processing environment using Go based concurrent workers to simulate coordinator participant communication and analyze commit behavior across multiple nodes.

- **Scalability Analysis Across Cluster Sizes:**  
Evaluated commit latency across clusters with 3, 5, 7, 9, and 11 nodes to study how batching influences scalability and transaction completion performance.

### Practical Significance and Impact
- **Lower Commit Latency:**
Batch based commit processing significantly decreases the time required to finalize transactions by reducing repeated coordination cycles that occur in conventional immediate commit protocols.

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






