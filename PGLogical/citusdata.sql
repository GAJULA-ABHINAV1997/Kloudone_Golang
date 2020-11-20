--Coordinator running on port:9700

CREATE EXTENSION citus;

--To connect coordinator node with worker1 node
SELECT * from master_add_node('localhost', 9701);
--To connect coordinator node with worker2 node
SELECT * from master_add_node('localhost', 9702);
--To verify that both worker nodes are connected to coordinator node
select * from master_get_active_worker_nodes();
 
CREATE TABLE companies (
    id bigint NOT NULL,
    name text NOT NULL,
    image_url text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE campaigns (
    id bigint NOT NULL,
    company_id bigint NOT NULL,
    name text NOT NULL,
    cost_model text NOT NULL,
    state text NOT NULL,
    monthly_budget bigint,
    blacklisted_site_urls text[],
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE ads (
    id bigint NOT NULL,
    company_id bigint NOT NULL,
    campaign_id bigint NOT NULL,
    name text NOT NULL,
    image_url text,
    target_url text,
    impressions_count bigint DEFAULT 0,
    clicks_count bigint DEFAULT 0,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

ALTER TABLE companies ADD PRIMARY KEY (id);
ALTER TABLE campaigns ADD PRIMARY KEY (id, company_id);
ALTER TABLE ads ADD PRIMARY KEY (id, company_id);

--citus replication model
SET citus.replication_model = 'streaming';

--distributing tables
SELECT create_distributed_table('companies', 'id');
SELECT create_distributed_table('campaigns', 'company_id');
SELECT create_distributed_table('ads', 'company_id');

--sample queries
INSERT INTO companies VALUES (5000, 'New Company', 'https://randomurl/image.png', now(), now());

UPDATE campaigns
SET monthly_budget = monthly_budget*2
WHERE company_id = 5;

SELECT name, cost_model, state, monthly_budget
FROM campaigns
WHERE company_id = 5
ORDER BY monthly_budget DESC
LIMIT 10;


--Worker1 node running on port:9701
 
CREATE EXTENSION citus;
--run this command to view the list of distributed tables
\d

--Worker2 node running on port:9702

CREATE EXTENSION citus;
--run this command to view the list of distributed tables
\d