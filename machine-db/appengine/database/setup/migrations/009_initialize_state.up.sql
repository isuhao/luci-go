-- Copyright 2018 The LUCI Authors.
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

ALTER TABLE datacenters ADD COLUMN state tinyint DEFAULT 0;
ALTER TABLE racks ADD COLUMN state tinyint DEFAULT 0;
ALTER TABLE switches ADD COLUMN state tinyint DEFAULT 0;
ALTER TABLE vlans ADD COLUMN state tinyint DEFAULT 0;
ALTER TABLE machines ADD COLUMN state tinyint DEFAULT 0;
ALTER TABLE physical_hosts ADD COLUMN state tinyint DEFAULT 0;
ALTER TABLE vms ADD COLUMN state tinyint DEFAULT 0;
