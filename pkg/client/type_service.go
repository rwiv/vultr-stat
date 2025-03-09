package client

type OsResponse struct {
	Os   []*OS `json:"os"`
	Meta *Meta `json:"meta"`
}

type OS struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`   // 풀네임
	Arch   string `json:"arch"`   // x32, x64...
	Family string `json:"family"` // windows, linux...
}

//export interface PlanResponse {
//	plans: Plan[];
//	meta: Meta;
//}
//
//// /plans
//export interface Plan {
//	id: string; // 줄임말 형태
//	vcpu_count: number; // cpu 코어 개수
//	ram: number; // 램 용량, MB 단위
//	disk: number; // 디스크 용량, GB 단위
//	disk_count: number; // 디스크 개수
//	bandwidth: number; // 대역폭, MB 단위
//	monthly_cost: number; // 월별 가격, 달러 단위
//	type: string;
//	locations: any[];
//}
//

//type RegionResponse struct {
//	Regions []*Region `json:"regions"`
//	Meta *Meta `json:"meta"`
//}

//// /regions
//export interface Region {
//	id: string; // 줄임말 형태
//	city: string; // 도시명
//	country: string; // 나라명 줄임말
//	continent: string; // 대륙명
//	options: any[];
//}
