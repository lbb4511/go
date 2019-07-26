window.$docsify = {
  name: "Golang 学习笔记",
  repo: "lbb4511/go",
  maxLevel: 2,
  subMaxLevel: 1,
  coverpage: true,
  onlyCover: true,
  loadSidebar: true,
  search: {
    maxAge: 86400000, // 过期时间，单位毫秒，默认一天
    placeholder: '搜索',
    noData: '找不到结果',
    depth: 2
  },
};

const gitalk = new Gitalk({
  clientID: '8194e2d7dba9b720d111',
  clientSecret: '801009aef54ad4261602b62d58ec1f4d6c571ae4',
  repo: 'go',
  owner: 'lbb4511',
  admin: ['lbb4511'],
  // facebook-like distraction free mode
  distractionFreeMode: true,
  perPage: 10

})