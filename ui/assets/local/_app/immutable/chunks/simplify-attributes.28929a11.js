import{f as r}from"./format-date.100c5f23.js";import{j as s}from"./format-time.66fea38b.js";const n=["closeTime","createTime","currentAttemptScheduledTime","earliestTime","eventTime","executionTime","expirationTime","expireTime","lastAccessTime","lastHeartbeatTime","lastStartedTime","lastUpdateTime","latestTime","releaseTime","scheduledTime","startedTime","startTime","workflowExecutionExpirationTime"],f=["backoffStartInterval","defaultWorkflowTaskTimeout","firstWorkflowTaskBackoff","heartbeatTimeout","initialInterval","maximumInterval","scheduleToCloseTimeout","scheduleToStartTimeout","startToCloseTimeout","startToFireTimeout","workflowExecutionRetentionPeriod","workflowExecutionRetentionTtl","workflowExecutionTimeout","workflowRunTimeout","workflowTaskTimeout"],a=e=>{for(const t of n)if(t===e)return!0;return!1},m=e=>{for(const t of f)if(t===e)return!0;return!1},l=e=>{if(e===null||e===void 0||typeof e!="object")return!1;const t=Object.keys(e),[o]=t;return!(t.length!==1||typeof e[o]!="string")},T=e=>{for(const t of Object.values(e))return t};function k(e,t=!1){for(const[o,i]of Object.entries(e))l(i)&&(e[o]=T(i)),a(o)&&!t&&(e[o]=r(i)),m(o)&&(e[o]=s(i));return e}export{k as s};
