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
        <head-bar v-if="!hide" :classList="classList" />
        <router-view />
        <foot-bar v-if="!hide" :info="user" />
        <n-back-top :visibility-height="370" :right="20" :bottom="20" />
      </div>
    </n-scrollbar>
  </n-config-provider>
</template>

<script>
import { zhCN, dateZhCN, darkTheme, lightTheme } from 'naive-ui'
import GlobalApi from './components/globalApi.vue'
import HeadBar from './components/headBar.vue'
import FootBar from './components/footBar.vue'
import { common, article } from './plugins/api'
import { useThemeStore } from './plugins/store'
import theme from './theme.json'

export default {
  name: "App",
  components: { GlobalApi, HeadBar },
  data: () => ({
    i18n: {
      main: zhCN,
      date: dateZhCN
    },
    hide: false,
    classList: [],
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
      this.initClass();
      this.initTags();
      this.initUser();
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
    },
    initClass() {
      article.getClass().then(res => {
        if (!res.state && res.message == '请完成初始化') {
          this.$router.replace('/install')
          this.hide = true
        }
        this.classList = res.state ? res.data : [];
        localStorage.setItem('cache:class', res.state ? JSON.stringify(res.data) : '[]')
      }).catch(err => {
        localStorage.setItem('cache:class', '[]')
      })
    },
    initTags() {
      article.getTags().then(res => {
        if (!res.state && res.message == '请完成初始化') {
          this.$router.replace('/install')
          this.hide = true
        }
        localStorage.setItem('cache:tags', res.state ? JSON.stringify(res.data) : '[]')
      }).catch(err => {
        localStorage.setItem('cache:class', '[]')
      })
    },
  },
  mounted() {
    document.body.classList.toggle('dark', this.isDark);
    this.init()
  }
};
</script>