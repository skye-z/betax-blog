<template>
    <n-drawer v-model:show="show" :width="370" placement="right">
        <n-drawer-content title="标签管理">
            <div v-if="loading" class="loading-block"><n-spin /></div>
            <n-input-group class="mb-10" v-for="(item, i) in list">
                <n-input v-model:value="item.name" type="text" placeholder="编辑时不能为空哦" @blur="e => update(e, i)" />
                <n-button type="error" round @click="remove(item.id)">删除</n-button>
            </n-input-group>
            <n-input-group class="mb-10">
                <n-input v-model:value="form" type="text" placeholder="编辑时不能为空哦" />
                <n-button type="success" round @click="add">新增</n-button>
            </n-input-group>
        </n-drawer-content>
    </n-drawer>
</template>
<script>
import { article, setting } from '../plugins/api'

export default {
    name: "TagsEdit",
    data: () => ({
        show: false,
        loading: false,
        cache: [],
        list: [],
        form: ''
    }),
    methods: {
        open() {
            this.loading = true
            this.init()
            this.show = true
        },
        init() {
            article.getTags().then(res => {
                if (res.state) {
                    this.list = res.data
                    this.cache = JSON.parse(JSON.stringify(res.data))
                } else window.$message.warning('列表加载失败');
                this.loading = false
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
                this.loading = false
            })
        },
        add() {
            if (this.form == '') {
                window.$message.warning('标签名不能为空');
                return false
            }
            setting.addTags(this.form).then(res => {
                if (res.state) {
                    this.init();
                    this.form = ''
                    window.$message.success('标签新增成功');
                } else window.$message.warning(res.message ? res.message:'新增标签失败');
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        update(e, i) {
            let value = e.target.value
            if (value == '') {
                window.$message.warning('标签名不能为空');
                return false
            }
            let check = this.cache[i]
            if (value == check.name) return false
            let item = this.list[i]
            setting.editTags(item.id, value).then(res => {
                if (res.state) {
                    this.cache = JSON.parse(JSON.stringify(this.list))
                    window.$message.success('标签更新成功');
                } else window.$message.warning(res.message ? res.message:'更新标签失败');
                
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        remove(id) {
            window.$dialog.warning({
                title: '操作确认',
                content: '删除标签后, 原归属于此标签的文章将自动清除标签关联, 确认要删除此标签吗?',
                positiveText: '确认删除',
                negativeText: '取消',
                onPositiveClick: () => {
                    setting.removeTags(id).then(res => {
                        if (res.state) {
                            this.init();
                            window.$message.success('标签删除成功');
                        }
                        else window.$message.warning('删除标签失败');
                    }).catch(err => {
                        window.$message.warning("发生意料之外的错误");
                    })
                }
            })
        }
    }
};
</script>
<style scoped></style>