import request from './request'

function post(url, data) {
    let cache = new URLSearchParams();
    for (let key in data) {
        cache.append(key, data[key]);
    }

    return request({
        url: url,
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        data: cache
    })
}

function get(url) {
    return request({
        url: url,
        method: 'GET'
    })
}

export const article = {
    getInfo: id => get('/article/details/' + id),
    search: keyword => post('/article/search', { keyword }),
    getNumber: (keyword, state) => post('/article/number', { keyword, state }),
    getList: (isBanner, isUp, state, page, number) => post('/article/list', { isBanner: isBanner ? 1:0, isUp: isUp ? 1:0, state, page, number }),
    getClass: () => get('/class'),
    getTags: () => get('/tags')
}