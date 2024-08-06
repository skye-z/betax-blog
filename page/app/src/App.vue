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
    <div id="app-center">
      <head-bar />
      <router-view />
      <foot-bar />
      <n-back-top :visibility-height="370" :right="20" :bottom="20" />
    </div>
  </n-config-provider>
</template>

<script>
import { zhCN, dateZhCN, darkTheme } from 'naive-ui'
import GlobalApi from './components/globalApi.vue'
import HeadBar from './components/headBar.vue'
import FootBar from './components/footBar.vue'
import theme from './theme.json'

export default {
  name: "App",
  components: { GlobalApi, HeadBar },
  data: () => ({
    theme: darkTheme,
    i18n: {
      main: zhCN,
      date: dateZhCN
    }
  }),
  methods: {
  }
};
</script>