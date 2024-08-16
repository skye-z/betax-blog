<template>
    <div class="app-content no-select">
        <div v-if="loading" class="loading-block"><n-spin /></div>
        <article-list :list="list" />
    </div>
</template>

<script>
import ArticleList from '../components/articleList.vue'
import { article } from '../plugins/api'

export default {
    name: "Any",
    components: { ArticleList },
    data: () => ({
        loading: true,
        list: [],
        classMapping: {},
        tags: [],
    }),
    methods: {
        init() {
            this.initData()
            this.loading = true
            this.getList()
        },
        initData() {
            this.list = []
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
        getList() {
            article.getAnyList().then(res => {
                if (res.state) {
                    if (res.data == null) res.data = []
                    this.next = res.data.length == this.number
                    if (res.data.length == 0) {
                        window.$message.warning('已经到底啦');
                        return false;
                    }
                    for (let i in res.data) {
                        this.list.push(res.data[i])
                    }
                    if (this.next) this.page = this.page + 1;
                } else {
                    window.$message.warning('获取文章列表失败');
                }
                this.loading = false;
            }).catch(err => {
                this.loading = false;
                window.$message.warning("获取文章列表出错");
            })
        },
        searchList() {
            article.search(this.keyword, this.page, this.number).then(res => {
                if (res.state) {
                    this.list = res.data;
                } else {
                    window.$message.warning('获取文章列表失败');
                }
                this.loading = false;
            }).catch(err => {
                this.loading = false;
                window.$message.warning("获取文章列表出错");
            })
        }
    },
    mounted() {
        this.init()
    },
    watch: {
        '$route': function (to, from) {
            if (from.name == to.name) {
                this.init()
            }
        },
    },
};
</script>

<style scoped>
.load-more {
    width: 40%;
}
</style>
