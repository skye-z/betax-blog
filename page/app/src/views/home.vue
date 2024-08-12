<template>
    <div class="app-content no-select">
        <div class="flex mb-10">
            <article-banner class="mr-10" :list="banner" />
            <div class="full-width">
                <div class="card readme mb-10 full-width">
                    <div class="flex pa-10 border-bottom">
                        <n-avatar class="mr-10" size="large"
                            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                        <div>
                            <div>Skye Zhang</div>
                            <div class="text-gray text-small">@skye-z</div>
                        </div>
                    </div>
                    <div>
                        <n-button class="social-btn" quaternary>
                            <template #icon>
                                <n-icon>
                                    <Github />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <QqOutlined />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <Linkedin />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <Discord />
                                </n-icon>
                            </template>
                        </n-button>
                    </div>
                </div>
                <div class="card tag-cloud pa-10 full-width">标签云</div>
            </div>
        </div>
        <div class="article-box">
            <article-list :list="tops" />
            <article-list :list="list" />
            <div class="text-center">
                <n-button class="load-more" :disabled="!next" strong secondary
                    @click="getList(false, false)">加载更多</n-button>
            </div>
        </div>
    </div>
</template>

<script>
import ArticleList from '../components/articleList.vue'
import ArticleBanner from '../components/articleBanner.vue'
import { Github, Linkedin, Discord } from '@vicons/fa'
import { QqOutlined } from '@vicons/antd'
import { article } from '../plugins/api'

export default {
    name: "Home",
    components: { Github, Linkedin, QqOutlined, Discord, ArticleList, ArticleBanner },
    data: () => ({
        loading: true,
        next: true,
        number: 20,
        page: 1,
        list: [],
        tops: [],
        banner: []
    }),
    methods: {
        init() {
            this.getList(true, false);
            this.getList(false, true);
            this.getList(false, false);
        },
        getList(isBanner, isUp) {
            article.getList(isBanner, isUp, -1, 2, this.page, this.number).then(res => {
                if (res.state) {
                    if (res.data == null) res.data = []
                    if (isBanner) this.banner = res.data ? res.data : []
                    else if (isUp) this.tops = res.data ? res.data : []
                    else {
                        let num = 0
                        let list = []
                        for (let i in res.data) {
                            let item = res.data[i]
                            if (item.isBanner || item.isUp) continue
                            list.push(item)
                            num++
                        }
                        this.next = num == this.number
                        if (res.data.length == 0) {
                            window.$message.warning('已经到底啦');
                            return false;
                        }
                        for (let i in res.data) {
                            this.list.push(res.data[i])
                        }
                        if (this.next) this.page = this.page + 1;
                    }
                    this.loading = false;
                } else {
                    window.$message.warning('获取文章列表失败');
                    this.loading = false;
                }
            }).catch(err => {
                this.loading = false;
                window.$message.warning("获取文章列表出错");
            })
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
.welcome {
    font-size: 32px;
    font-weight: bold;
}

.readme {
    height: 97px;
}

.tag-cloud {
    height: 192px;
}

.article-list {
    height: 200px;
}

.load-more {
    width: 40%;
}

.readme:deep(.social-btn) {
    width: 25%;
}
</style>
