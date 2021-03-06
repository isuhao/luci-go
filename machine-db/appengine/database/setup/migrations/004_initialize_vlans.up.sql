-- Copyright 2017 The LUCI Authors.
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--      http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

-- Required fields are not enforced by this schema.
-- The Machine Database will enforce any such constraints.

CREATE TABLE IF NOT EXISTS vlans (
	-- The ID of this VLAN.
	id int,
	-- An alias for this VLAN.
	alias varchar(255),
	-- The CIDR block of IPv4 addresses belonging to this VLAN.
	cidr_block varchar(255),
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS ips (
	id int NOT NULL AUTO_INCREMENT,
	-- The IPv4 address of this IP.
	ipv4 int unsigned,
	-- The VLAN this IP belongs to.
	vlan_id int,
	PRIMARY KEY (id),
	FOREIGN KEY (vlan_id) REFERENCES vlans (id) ON DELETE CASCADE,
	UNIQUE (ipv4)
);
