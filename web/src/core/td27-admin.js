// 加载网站配置文件夹
import { register } from './global'

export default {
    install: (app) => {
        register(app)
        console.log(`
       欢迎使用 TD27-Admin-Template
       当前版本:v1.0.1
    `)
    }
}
