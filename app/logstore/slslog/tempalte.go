package slslogComponent

// 访问情况 统计  pv uv
const PvUvTotal = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
and category :"view" |
SELECT
  COUNT(*) as pv,
  approx_distinct("_uuid") as uv`

// 访问情况 时间线 pv uv
const PvUvTrend = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
and category :"view" |
SELECT
  approx_distinct("_uuid") as uv,
  COUNT(*) as pv,
  time_series(
    __time__,
    '2h',
    '%Y-%m-%d %H:%i:%S',
    '0'
  ) as ts
GROUP by
  ts
ORDER by
  ts`

// 访问情况 sdk 版本
const SdkVersionCount = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
SELECT
  _sdk_v as version,
  COUNT(*) as count
GROUP by
  version
ORDER BY
  count DESC`

// 访问情况 type count
const CategoryCount = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
SELECT
  COUNT(*) as count,
  category as category
GROUP BY
  category
ORDER BY
  count DESC `

// 访问情况 入口页面 页面访问情况
const PageTotal = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: "view" |
SELECT
  "_url" as url,
  COUNT(*) as pv,
  approx_distinct("_uuid") as uv
GROUP BY
  url
ORDER BY
  pv DESC
LIMIT
  20`

// 访问情况 城市分布
const CityPv = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: "view" |
SELECT
  COUNT(*) as pv,
  ip_to_city(ip) as city
GROUP by
  city
ORDER BY
  pv DESC
LIMIT
  20`

// 异常 统计 error 总量
const ErrorCount = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
and category :"error" |
SELECT
  COUNT(*) as count,
  approx_distinct("_uuid") as effect_user`

// 异常 （错误量） 及（影响用户量）
const ErrorCountTrend = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
and category :"error" |
SELECT
  COUNT(*) as count,
  approx_distinct("_uuid") as effect_user,
  time_series(
    __time__,
    '2h',
    '%Y-%m-%d %H:%i:%S',
    '0'
  ) as ts
GROUP by
  ts
ORDER by
  ts`

// api
const ApiErrorCount = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
and category :"api" |
SELECT
  COUNT(*) as count,
  approx_distinct("_uuid") as effect_user`

// api 接口错误趋势
const ApiErrorTrend = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
and category :"api" |
SELECT
  COUNT(*) as count,
  approx_distinct("_uuid") as effect_user,
  time_series(
    __time__,
    '2h',
    '%Y-%m-%d %H:%i:%S',
    '0'
  ) as ts
GROUP by
  ts
ORDER by
  ts`

// api 接口错误排行列表
const ApiErrorList = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
and category :"api" |
SELECT
  url,
  method,
  error_type,
  COUNT(*) as count,
  approx_distinct(_uuid) as effect_user
GROUP BY
  url,
  method,
  error_type
ORDER BY
  count DESC`

// 性能 页面打开
const PerfNavigationTimingTrend = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: view
AND type: navigationTiming
AND totalTime > 0 |
SELECT
  round(avg(dnsLookupTime), 2) as dnsLookupTime,
  round(avg(downloadTime), 2) as downloadTime,
  round(avg(fetchTime), 2) as fetchTime,
  round(avg(headerSize), 2) as headerSize,
  round(avg(timeToFirstByte), 2) as timeToFirstByte,
  round(avg(totalTime), 2) as totalTime,
  time_series(
    __time__,
    '2h',
    '%Y-%m-%d %H:%i:%S',
    '0'
  ) as ts
GROUP by
  ts
ORDER by
  ts`

const PerfNavigationTimingValues = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: view
AND type: navigationTiming 
AND totalTime > 0 |
SELECT
  round(avg(dnsLookupTime), 2) as dnsLookupTime,
  round(avg(downloadTime), 2) as downloadTime,
  round(avg(fetchTime), 2) as fetchTime,
  round(avg(headerSize), 2) as headerSize,
  round(avg(timeToFirstByte), 2) as timeToFirstByte,
  round(avg(totalTime), 2) as totalTime`

// 性能 资源加载
const PerfDataConsumption = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: perf
AND type: dataConsumption 
AND total > 0 |
SELECT
  round(avg(css),2) as css,
  round(avg(img),2) as img,
  round(avg(other),2) as other,
  round(avg(script),2) as script,
  round(avg(total),2) as total,
  round(avg(xmlhttprequest),2) as xhr,
  round(avg(fetch),2) as fetch,
  time_series(
    __time__,
    '2h',
    '%Y-%m-%d %H:%i:%S',
    '0'
  ) as ts
GROUP by
  ts
ORDER by
  ts`

const PerfDataConsumptionValues = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: perf
AND type: dataConsumption 
AND total > 0 |
SELECT
  round(avg(css),2) as css,
  round(avg(img),2) as img,
  round(avg(other),2) as other,
  round(avg(script),2) as script,
  round(avg(total),2) as total,
  round(avg(xmlhttprequest),2) as xhr,
  round(avg(fetch),2) as fetch`

// 性能 页面
const PerfMetrics = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: perf
AND type: metrics 
AND fp > 0 |
SELECT
  round(avg(fp),2) as fp,
  round(avg(fcp),2) as fcp,
  round(avg(lcp),2) as lcp,
  round(avg(fid),2) as fid,
  round(avg(cls),2) as cls,
  round(avg(tbt),2) as tbt,
  time_series(
    __time__,
    '2h',
    '%Y-%m-%d %H:%i:%S',
    '0'
  ) as ts
GROUP by
  ts
ORDER by
  ts`

const PerfMetricsValues = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: perf
AND type: metrics 
AND fp > 0 |
SELECT
  round(avg(fp),2) as fp,
  round(avg(fcp),2) as fcp,
  round(avg(lcp),2) as lcp,
  round(avg(fid),2) as fid,
  round(avg(cls),2) as cls,
  round(avg(tbt),2) as tbt`

// 资源加载 失败趋势图
const ResLoadFailTotalTrend = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: res |
SELECT
  COUNT(*) as count,
  approx_distinct("_uuid") as effect_user,
  time_series(
    __time__,
    '2h',
    '%Y-%m-%d %H:%i:%S',
    '0'
  ) as ts
GROUP by
  ts
ORDER by
  ts`

// 资源加载 失败总数
const ResLoadFailTotal = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: res |
SELECT
  COUNT(*) as count,
  approx_distinct("_uuid") as effect_user`

// 资源加载 失败排行
const ResLoadFailList = `*
AND _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5
AND category: res |
SELECT
  src,
  COUNT(*) as count,
  approx_distinct("_uuid") as effect_user
GROUP BY
  src
ORDER BY
  count
LIMIT
  20`

// Ip
const ProjectIpToCountry = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  ip_to_country(ip) as ip_country,
  count(*) as pv,
  approx_distinct(_uuid) as uv
WHERE
  ip_to_domain(ip) != 'intranet'
group by
  ip_country
order by
  pv desc
limit
  50`

const ProjectIpToProvince = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  ip_to_province(ip) as province,
  count(*) as pv,
  approx_distinct(_uuid) as uv
WHERE
  ip_to_domain(ip) != 'intranet'
group by
  province
order by
  pv desc`

const ProjectIpToCity = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  ip_to_city(ip) as city,
  count(*) as pv,
  approx_distinct(_uuid) as uv
WHERE
  ip_to_domain(ip) != 'intranet'
group by
  city
order by
  pv desc`

//
const ProjectEventCount = `* and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 | SELECT COUNT(*) as count`

const ProjectSendMode = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  _send_mode as mode,
  COUNT(*) as count
GROUP BY
  mode
ORDER BY
  count DESC`

const ProjectEnv = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  _env as env,
  COUNT(*) as count
GROUP BY
  env
ORDER BY
  count DESC`

const ProjectVersion = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  _version as version,
  COUNT(*) as count
GROUP BY
  version
ORDER BY
  count DESC`

const projectUserScreen = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  _s_wh as wh,
  COUNT(*) as count
GROUP BY
  wh
ORDER BY
  count DESC`

const projectCategory = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  category as category,
  COUNT(*) as count
GROUP BY
  category
ORDER BY
  count DESC`

const projectType = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  type as type,
  COUNT(*) as count
GROUP BY
  type
ORDER BY
  count DESC`

const projectBrowser = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  _ua_browser as browser,
  COUNT(*) as count,
  approx_distinct(_uuid) as user
GROUP BY
  browser
ORDER BY
  user DESC`

const projectEngine = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  _ua_engine as engine,
  COUNT(*) as count,
  approx_distinct(_uuid) as user
GROUP BY
  engine
ORDER BY
  user DESC`

const projectOs = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  _ua_os as os,
  COUNT(*) as count,
  approx_distinct(_uuid) as user
GROUP BY
  os
ORDER BY
  user DESC`

const projectPlatform = `*
and _appId: fca5deec-a9db-4dac-a4db-b0f4610d16a5 |
select
  _ua_platform as platform,
  COUNT(*) as count,
  approx_distinct(_uuid) as user
GROUP BY
  platform
ORDER BY
  user DESC`
