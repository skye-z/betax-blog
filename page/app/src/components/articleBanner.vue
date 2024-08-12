<template>
    <div class="card carousel">
        <n-carousel direction="vertical" dot-placement="right" show-arrow style="border-radius: 8px;">
            <div v-if="list.length == 0" class="card carousel-item welcome flex align-center justify-center">
                <div>欢迎来访 👏</div>
            </div>
            <div v-for="item in list" class="card carousel-item pa-10" @click="open(item.id)">
                <div class="banner-title">{{ item.title }}</div>
                <n-time v-if="item.releaseTime" class="text-gray" :time="item.releaseTime" />
                <div class="banner-abstract">{{ item.abstract }}</div>
                <div class="banner-foot flex align-center">
                    <n-tag class="mr-5" size="small" :bordered="false" type="success">
                        {{ classMapping[item.class] }}
                    </n-tag>
                    <template v-if="item.tags">
                        <n-tag class="mr-5" size="small" v-for="tag in item.tags" :bordered="false" type="info">
                            {{ tag.name }}
                        </n-tag>
                    </template>
                </div>
            </div>
        </n-carousel>
    </div>
</template>
<script>
export default {
    name: "ArticleBanner",
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
.carousel {
    width: 580px;
    height: 300px;
}

.carousel-item {
    width: 580px;
    height: 300px;
    position: relative;
    cursor: pointer;
}

.banner-title {
    font-size: 24px;
    font-weight: bold;
}

.banner-abstract {
    font-size: 18px;
    line-height: 1.5;
    width: calc(100% - 20px);

    display: -webkit-box;
    -webkit-line-clamp: 6;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
}

.banner-foot {
    position: absolute;
    bottom: 10px;
    left: 10px;
}
</style>