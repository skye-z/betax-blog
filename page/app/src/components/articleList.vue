<template>
    <div v-for="item in list" class="card article-item mb-10 pa-10" @click="open(item.id)">
        <div class="mb-5 title">{{ item.title }}</div>
        <div class="flex align-center">
            <n-tag class="mr-5" v-if="item.releaseTime" size="small" :bordered="false">
                <n-time :time="item.releaseTime" type="relative" />
            </n-tag>
            <n-tag class="mr-5" size="small" :bordered="false" type="success">
                {{ classMapping[item.class] }}
            </n-tag>
            <template v-if="item.tags">
                <n-tag class="mr-5" size="small" v-for="tag in item.tags" :bordered="false" type="info">
                    {{ tag.name }}
                </n-tag>
            </template>
            <n-tag class="mr-5" size="small" v-if="item.isUp" :bordered="false" type="error">置顶</n-tag>
        </div>
        <div v-if="item.abstract" class="mt-10 pa-10 abstract">{{ item.abstract }}</div>
    </div>
</template>
<script>
export default {
    name: "ArticleList",
    props: {
        list: {
            type: Array,
            default: () => []
        }
    },
    data: () => ({
        classMapping: {},
        tags: [],
    }),
    methods: {
        init() {
            let cacheTags = localStorage.getItem('cache:tags');
            if (cacheTags != undefined) {
                this.tops = JSON.parse(cacheTags)
            }
            let cacheClass = localStorage.getItem('cache:class');
            if (cacheClass != undefined) {
                cacheClass = JSON.parse(cacheClass)
                for (let i in cacheClass) {
                    this.classMapping[cacheClass[i].id] = cacheClass[i].name
                }
            }
        },
        open(id) {
            this.$router.push('/info/' + id)
        }
    },
    mounted() {
        this.init()
    },
};
</script>
<style scoped>
.article-item {
    cursor: pointer;
}

.article-item:hover {
    background-color: var(--color-light-3);
}

.dark .article-item:hover {
    background-color: var(--color-dark-3);
}

.title {
    font-size: 16px;
}

.abstract {
    background-color: var(--color-light-1);
}

.dark .abstract {
    background-color: var(--color-dark-3);
}
</style>