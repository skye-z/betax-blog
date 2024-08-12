<template>
    <div class="head-bar pa-10 flex align-center justify-between">
        <div class="flex align-center">
            <img class="logo mr-10" src="../assets/logo.png" />
            <n-button quaternary class="mr-10" @click="jump('/')">首页</n-button>
            <n-button quaternary class="mr-10" @click="jump('/list')">随便看看</n-button>
            <n-dropdown trigger="hover" :options="classList" key-field="id" label-field="name"  @select="selectClass">
                <n-button quaternary>
                    <template #icon>
                        <n-icon>
                            <MoreHorizontal32Filled />
                        </n-icon>
                    </template>
                </n-button>
            </n-dropdown>
        </div>
        <n-input class="search" v-model:value="keyword" round placeholder="搜索文章" @keyup.enter="search">
            <template #prefix>
                <n-icon>
                    <Search24Filled />
                </n-icon>
            </template>
        </n-input>
        <div class="right-bar flex align-center justify-end">
            <n-button quaternary circle>
                <template #icon>
                    <n-icon>
                        <Lightbulb24Filled />
                    </n-icon>
                </template>
            </n-button>
        </div>
    </div>
</template>

<script>
import { MoreHorizontal32Filled, Search24Filled, Lightbulb24Filled } from '@vicons/fluent'

export default {
    name: "HeadBar",
    components: { MoreHorizontal32Filled, Search24Filled, Lightbulb24Filled },
    props: {
        classList: {
            type: Array,
            default: () => []
        }
    },
    data: () => ({
        keyword: ''
    }),
    methods: {
        jump(path) {
            this.$router.push(path)
        },
        selectClass(e){
            this.jump('/list?class='+e)
        },
        search(){
            if(this.keyword == '') return false
            this.jump('/list?q='+this.keyword)
            this.keyword = ''
        }
    }
};
</script>

<style scoped>
.head-bar {
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

.right-bar{
    width: 260px;
}
</style>
