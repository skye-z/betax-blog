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
        <head-bar />
        <router-view />
        <foot-bar :info="user" />
        <n-back-top :visibility-height="370" :right="20" :bottom="20" />
      </div>
    </n-scrollbar>
  </n-config-provider>
</template>

<script>
import { zhCN, dateZhCN, darkTheme, lightTheme } from 'naive-ui'
import GlobalApi from './components/globalApi.vue'
import { useThemeStore } from './plugins/store'
import HeadBar from './components/headBar.vue'
import FootBar from './components/footBar.vue'
import theme from './theme.json'
import { common } from './plugins/api'


export default {
  name: "App",
  components: { GlobalApi, HeadBar },
  data: () => ({
    i18n: {
      main: zhCN,
      date: dateZhCN
    },
    user: {}
  }),
  computed: {
    isDark() {
      const themeStore = useThemeStore();
      return themeStore.isDark;
    },
    theme() {
      const themeStore = useThemeStore();
      return themeStore.isDark ? darkTheme : lightTheme;
    }
  },
  methods: {
    init() {
      this.initUser();
      common.getPing().then(res => {
        console.log((res.time * 1000) - new Date().getTime())
      })
    },
    initUser() {
      common.getUser().then(res => {
        if (res.state) {
          if (res.data.bio.length > 70) res.data.bio = res.data.bio.substring(0, 68) + '...'
          this.user = res.data
        }
        localStorage.setItem('cache:user', res.state ? JSON.stringify(res.data) : '{}')
      }).catch(err => {
        localStorage.setItem('cache:user', '{}')
      })
    }
  },
  mounted() {
    document.body.classList.toggle('dark', this.isDark);
    this.init()
  }
};
</script>