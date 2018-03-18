declare module '*.vue' {
    import Vue from 'vue'
    export default typeof Vue
}

declare module 'nuxt' {
    export const Nuxt: any;
    export const Builder: any;
}
