export const policyOptions = [
  'client-readwrite',
  'client-readonly',
  'network',
  'monitoring',
  'internal'
].map(val => ({label: val, value: val}))
