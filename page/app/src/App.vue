<template>
  <n-config-provider :locale="i18n.main" :date-locale="i18n.date" :theme-overrides="theme">
    <n-loading-bar-provider>
      <n-dialog-provider>
        <n-message-provider>
          <n-notification-provider>
            <global-api />
          </n-notification-provider>
        </n-message-provider>
      </n-dialog-provider>
    </n-loading-bar-provider>
    <n-scrollbar style="height: 100vh">
      <div id="app-center">
        <head-bar :classList="classList" />
        <router-view />
        <foot-bar />
        <n-back-top :visibility-height="370" :right="20" :bottom="20" />
      </div>
    </n-scrollbar>
  </n-config-provider>
</template>

<script>
import { zhCN, dateZhCN, darkTheme } from 'naive-ui'
import GlobalApi from './components/globalApi.vue'
import HeadBar from './components/headBar.vue'
import FootBar from './components/footBar.vue'
import { article } from './plugins/api'
import theme from './theme.json'

export default {
  name: "App",
  components: { GlobalApi, HeadBar },
  data: () => ({
    theme: darkTheme,
    i18n: {
      main: zhCN,
      date: dateZhCN
    },
    classList: []
  }),
  methods: {
    init() {
      this.initClass();
      this.initTags();
    },
    initClass() {
      article.getClass().then(res => {
        this.classList = res.state ? res.data:[];
        localStorage.setItem('cache:class', res.state ? JSON.stringify(res.data) : '[]')
      }).catch(err => {
        localStorage.setItem('cache:class', '[]')
      })
    },
    initTags() {
      article.getTags().then(res => {
        localStorage.setItem('cache:tags', res.state ? JSON.stringify(res.data) : '[]')
      }).catch(err => {
        localStorage.setItem('cache:class', '[]')
      })
    },
  },
  mounted() {
    this.init()
  }
};
</script>