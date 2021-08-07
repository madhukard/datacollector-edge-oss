// Copyright 2018 StreamSets Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package origins

import (
	_ "github.com/streamsets/datacollector-edge/stages/origins/dev_data_generator"
	_ "github.com/streamsets/datacollector-edge/stages/origins/dev_random"
	_ "github.com/streamsets/datacollector-edge/stages/origins/dev_rawdata"
	_ "github.com/streamsets/datacollector-edge/stages/origins/filetail"
	_ "github.com/streamsets/datacollector-edge/stages/origins/httpclient"
	_ "github.com/streamsets/datacollector-edge/stages/origins/httpserver"
	_ "github.com/streamsets/datacollector-edge/stages/origins/mqtt"
	_ "github.com/streamsets/datacollector-edge/stages/origins/sensor_reader"
	_ "github.com/streamsets/datacollector-edge/stages/origins/spooler"
	_ "github.com/streamsets/datacollector-edge/stages/origins/system_metrics"
	_ "github.com/streamsets/datacollector-edge/stages/origins/websocketclient"
	_ "github.com/streamsets/datacollector-edge/stages/origins/windows"
)
