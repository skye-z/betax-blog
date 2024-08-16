<template>
    <n-drawer v-model:show="show" :width="370" placement="right">
        <n-drawer-content title="分类管理">
            <div v-if="loading" class="loading-block"><n-spin /></div>
            <n-input-group class="mb-10" v-for="(item, i) in list">
                <n-input v-model:value="item.name" type="text" placeholder="编辑时不能为空哦" @blur="e => update(e, i)" />
                <n-button type="error" round @click="remove(item.id)">删除</n-button>
            </n-input-group>
            <n-input-group class="mb-10">
                <n-input v-model:value="form.name" type="text" placeholder="编辑时不能为空哦" />
                <n-button type="success" round @click="add">新增</n-button>
            </n-input-group>
        </n-drawer-content>
    </n-drawer>
</template>
<script>
import { article, setting } from '../plugins/api'

export default {
    name: "ClassEdit",
    data: () => ({
        show: false,
        loading: false,
        cache: [],
        list: [],
        form: {
            name: '',
            superior: -1
        }
    }),
    methods: {
        open() {
            this.loading = true
            this.init()
            this.show = true
        },
        init() {
            article.getClass().then(res => {
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
            if (this.form.name == '') {
                window.$message.warning('分类名不能为空');
                return false
            }
            setting.addClass(this.form.name, this.form.superior).then(res => {
                if (res.state) {
                    this.init();
                    this.form.name = ''
                    window.$message.success('分类新增成功');
                } else window.$message.warning(res.message ? res.message:'新增分类失败');
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        update(e, i) {
            let value = e.target.value
            if (value == '') {
                window.$message.warning('分类名不能为空');
                return false
            }
            let check = this.cache[i]
            if (value == check.name) return false
            let item = this.list[i]
            setting.editClass(item.id, value, item.superior).then(res => {
                if (res.state) {
                    this.cache = JSON.parse(JSON.stringify(this.list))
                    window.$message.success('分类更新成功');
                } else window.$message.warning(res.message ? res.message:'更新分类失败');
                
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        remove(id) {
            window.$dialog.warning({
                title: '操作确认',
                content: '删除分类后, 原归属于此分类的文章将自动清除分类关联, 确认要删除此分类吗?',
                positiveText: '确认删除',
                negativeText: '取消',
                onPositiveClick: () => {
                    setting.removeClass(id).then(res => {
                        if (res.state) {
                            this.init();
                            window.$message.success('分类删除成功');
                        }
                        else window.$message.warning('删除分类失败');
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