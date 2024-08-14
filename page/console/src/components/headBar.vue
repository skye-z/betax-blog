<template>
    <div class="head-bar pa-10 flex align-center justify-between">
        <div class="flex align-center">
            <img class="logo mr-10" src="../assets/logo.png" />
            <n-button quaternary class="mr-10" @click="jump('/')">首页</n-button>
            <n-button quaternary class="mr-10" @click="jump('/list')">文章</n-button>
            <n-dropdown trigger="hover" :options="class" @select="handleSelect">
                <n-button quaternary>
                    <template #icon>
                        <n-icon>
                            <MoreHorizontal32Filled />
                        </n-icon>
                    </template>
                </n-button>
            </n-dropdown>
        </div>
        <div class="right-bar flex align-center justify-end">
            <n-button quaternary circle class="mr-10">
                <template #icon>
                    <n-icon>
                        <SettingFilled />
                    </n-icon>
                </template>
            </n-button>
            <n-button quaternary circle @click="toggleTheme">
                <template #icon>
                    <n-icon>
                        <Lightbulb24Filled v-if="isDark" />
                        <LightbulbFilament24Filled v-else />
                    </n-icon>
                </template>
            </n-button>
        </div>
    </div>
</template>

<script>
import { MoreHorizontal32Filled, Lightbulb24Filled, LightbulbFilament24Filled } from '@vicons/fluent'
import { useThemeStore } from '../plugins/store'
import { SettingFilled } from '@vicons/antd'

export default {
    name: "HeadBar",
    components: { MoreHorizontal32Filled, SettingFilled, Lightbulb24Filled, LightbulbFilament24Filled },
    data: () => ({
        class: [{
            label: '分类管理',
            key: 'class'
        }, {
            label: '标签管理',
            key: 'tags'
        }, {
            label: '文件管理',
            key: 'file'
        }]
    }),
    computed: {
        isDark() {
            const themeStore = useThemeStore();
            return themeStore.isDark;
        }
    },
    methods: {
        jump(path) {
            this.$router.push(path)
        },
        toggleTheme() {
            const themeStore = useThemeStore();
            themeStore.toggleTheme();
        }
    }
};
</script>

<style scoped>
.head-bar {
    color: var(--text-color-light-2);
}

.dark .head-bar{
    color: var(--text-color-dark-2);
}

.logo {
    width: 48px;
}

.search {
    text-align: center;
    min-width: 200px;
    width: 30%;
}

.right-bar {
    width: 260px;
}
</style>
