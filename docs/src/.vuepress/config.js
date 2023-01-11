const { description } = require('../../package')

module.exports = {
    base: '/',
    /**
     * Ref：https://v1.vuepress.vuejs.org/config/#title
     */
    title: 'FYX Gaming - Fyx VM Documentation',
    /**
     * Ref：https://v1.vuepress.vuejs.org/config/#description
     */
    description: description,

    /**
     * Extra tags to be injected to the page HTML `<head>`
     *
     * ref：https://v1.vuepress.vuejs.org/config/#head
     */
    head: [
        ['meta', { name: 'theme-color', content: '#3eaf7c' }],
        ['meta', { name: 'apple-mobile-web-app-capable', content: 'yes' }],
        ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }]
    ],

    /**
     * Theme configuration, here is the default theme configuration for VuePress.
     *
     * ref：https://v1.vuepress.vuejs.org/theme/default-theme-config.html
     */
    themeConfig: {
        repo: '',
        logo: 'fyxlogo.png',
        editLinks: false,
        docsDir: '',
        editLinkText: '',
        lastUpdated: false,
        sidebar: [
            {
                title: 'Introduction',   // required
                path: '',      // optional, link of the title, which should be an absolute path and must exist
                collapsable: false, // optional, defaults to true
                sidebarDepth: 1,    // optional, defaults to 1
                children: [
                    '/'
                ]
            },
            {
                title: 'Getting Started',   // required
                path: '',      // optional, link of the title, which should be an absolute path and must exist
                collapsable: false, // optional, defaults to true
                sidebarDepth: 1,    // optional, defaults to 1
                children: [
                    '/subdocs/getting-started'
                ]
            },
            {
                title: 'Jigs',
                path: '',
                collapsable: false, // optional, defaults to true
                children: [
                    '/subdocs/jigs'
                ]
            },
            {
                title: 'Code',
                path: '',
                collapsable: false, // optional, defaults to true
                children: [
                    '/subdocs/code'
                ]
            },
            {
                title: 'Tokens',
                path: '',
                collapsable: false, // optional, defaults to true
                children: [
                    '/subdocs/tokens'
                ]
            },
            {
                title: 'SDK',
                path: '',
                collapsable: false, // optional, defaults to true
                children: [
                    '/subdocs/sdk'
                ]
            }
        ]
    },

    /**
     * Apply plugins，ref：https://v1.vuepress.vuejs.org/zh/plugin/
     */
    plugins: [
        '@vuepress/plugin-back-to-top',
        '@vuepress/plugin-medium-zoom',
    ]
}