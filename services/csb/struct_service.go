package csb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Service is a nested struct in csb response
type Service struct {
	ProvideType            string           `json:"ProvideType" xml:"ProvideType"`
	RouteConfJson          string           `json:"RouteConfJson" xml:"RouteConfJson"`
	ServiceProviderType    string           `json:"ServiceProviderType" xml:"ServiceProviderType"`
	StatisticName          string           `json:"StatisticName" xml:"StatisticName"`
	CreateTime             int              `json:"CreateTime" xml:"CreateTime"`
	AllVisiable            bool             `json:"AllVisiable" xml:"AllVisiable"`
	ServiceOpenRestfulPath string           `json:"ServiceOpenRestfulPath" xml:"ServiceOpenRestfulPath"`
	ConsumeTypesJSON       string           `json:"ConsumeTypesJSON" xml:"ConsumeTypesJSON"`
	SSL                    bool             `json:"SSL" xml:"SSL"`
	CasTargets             string           `json:"CasTargets" xml:"CasTargets"`
	ProjectName            string           `json:"ProjectName" xml:"ProjectName"`
	UserId                 string           `json:"UserId" xml:"UserId"`
	IpWhiteStr             string           `json:"IpWhiteStr" xml:"IpWhiteStr"`
	OldVersion             string           `json:"OldVersion" xml:"OldVersion"`
	ApproveUserId          string           `json:"ApproveUserId" xml:"ApproveUserId"`
	ModelVersion           string           `json:"ModelVersion" xml:"ModelVersion"`
	Scope                  string           `json:"Scope" xml:"Scope"`
	ServiceName            string           `json:"ServiceName" xml:"ServiceName"`
	Description            string           `json:"Description" xml:"Description"`
	Status                 int              `json:"Status" xml:"Status"`
	IpBlackStr             string           `json:"IpBlackStr" xml:"IpBlackStr"`
	ErrDefJSON             string           `json:"ErrDefJSON" xml:"ErrDefJSON"`
	PrincipalName          string           `json:"PrincipalName" xml:"PrincipalName"`
	Active                 bool             `json:"Active" xml:"Active"`
	OrderInfo              string           `json:"OrderInfo" xml:"OrderInfo"`
	Id                     int              `json:"Id" xml:"Id"`
	OttFlag                bool             `json:"OttFlag" xml:"OttFlag"`
	AccessParamsJSON       string           `json:"AccessParamsJSON" xml:"AccessParamsJSON"`
	PolicyHandler          string           `json:"PolicyHandler" xml:"PolicyHandler"`
	ModifiedTime           int              `json:"ModifiedTime" xml:"ModifiedTime"`
	OpenRestfulPath        string           `json:"OpenRestfulPath" xml:"OpenRestfulPath"`
	InterfaceName          string           `json:"InterfaceName" xml:"InterfaceName"`
	OwnerId                string           `json:"OwnerId" xml:"OwnerId"`
	ValidConsumeTypes      bool             `json:"ValidConsumeTypes" xml:"ValidConsumeTypes"`
	Qps                    int              `json:"Qps" xml:"Qps"`
	CsbId                  int              `json:"CsbId" xml:"CsbId"`
	Alias                  string           `json:"Alias" xml:"Alias"`
	ProjectId              int              `json:"ProjectId" xml:"ProjectId"`
	ValidProvideType       bool             `json:"ValidProvideType" xml:"ValidProvideType"`
	SkipAuth               bool             `json:"SkipAuth" xml:"SkipAuth"`
	ServiceVersion         string           `json:"ServiceVersion" xml:"ServiceVersion"`
	CasServTargets         []string         `json:"CasServTargets" xml:"CasServTargets"`
	ConsumeTypes           []string         `json:"ConsumeTypes" xml:"ConsumeTypes"`
	RouteConf              RouteConf        `json:"RouteConf" xml:"RouteConf"`
	ServiceVersionsList    []ServiceVersion `json:"ServiceVersionsList" xml:"ServiceVersionsList"`
	VisiableGroupList      []VisiableGroup  `json:"VisiableGroupList" xml:"VisiableGroupList"`
}
