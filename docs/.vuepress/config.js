module.exports = {
  theme: 'cosmos',
  title: 'DappLink Documentation',
  locales: {
    '/': {
      lang: 'en-US'
    },
  },
  markdown: {
    extendMarkdown: (md) => {
      md.use(require("markdown-it-katex"));
    },
  },
  head: [
    [
      "link",
      {
        rel: "stylesheet",
        href:
          "https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.5.1/katex.min.css",
      },
    ],
    [
      "link",
      {
        rel: "stylesheet",
        href:
          "https://cdn.jsdelivr.net/github-markdown-css/2.2.1/github-markdown.css",
      },
    ],
  ],
  base: process.env.VUEPRESS_BASE || '/',
  plugins: [
    'vuepress-plugin-element-tabs'
  ],
  head: [
    // ['link', { rel: "apple-touch-icon", sizes: "180x180", href: "/apple-touch-icon.png" }],
    ['link', { rel: "icon", type: "image/png", sizes: "32x32", href: "/favicon32.png" }],
    ['link', { rel: "icon", type: "image/png", sizes: "16x16", href: "/favicon16.png" }],
    ['link', { rel: "manifest", href: "/site.webmanifest" }],
    ['meta', { name: "msapplication-TileColor", content: "#2e3148" }],
    ['meta', { name: "theme-color", content: "#ffffff" }],
    ['link', { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" }],
    // ['link', { rel: "apple-touch-icon-precomposed", href: "/apple-touch-icon-precomposed.png" }],
  ],
  themeConfig: {
    repo: 'tharsis/dapplink',
    docsRepo: 'tharsis/dapplink',
    docsBranch: 'main',
    docsDir: 'docs',
    editLinks: true,
    custom: true,
    project: {
      name: 'DappLink',
      denom: 'DappLink',
      ticker: 'DappLink',
      binary: 'dapplinkd',
      testnet_denom: 'tDappLink',
      testnet_ticker: 'tDappLink',
      rpc_url: 'https://eth.bd.dapplink.org:8545',
      rpc_url_testnet: 'https://eth.bd.dapplink.dev:8545',
      rpc_url_local: 'http://localhost:8545/',
      chain_id: '9001',
      testnet_chain_id: '9000',
      latest_version: 'v5.0.0',
      version_number: '2',
      testnet_version_number: '4',
      testnet_evm_explorer_url: 'https://evm.dapplink.dev',
      evm_explorer_url: 'https://evm.dapplink.org',
      testnet_cosmos_explorer_url: 'https://explorer.dapplink.dev/',
      cosmos_explorer_url: 'https://www.mintscan.io/dapplink',
    },
    logo: {
      src: '/dapplink-black.svg',
    },
    algolia: {
      id: 'K3VQTEW3G5',
      key: 'bf836a3c934b1d4df091d5c5b69c65d7',
      index: 'dapplink'
    },
    topbar: {
      banner: false
    },
    sidebar: {
      auto: false,
      nav: [
        {
          title: 'About dapplink',
          children: [
            {
              title: 'Introduction',
              directory: true,
              path: '/about/intro'
            },
            {
              title: 'DappLink Ecosystem',
              path: 'https://dapplink.space/'
            },
            {
              title: 'Awesome dapplink',
              path: 'https://github.com/tharsis/awesome'
            },
          ]
        },
        {
          title: 'For Users',
          children: [
            {
              title: 'Basic Concepts',
              directory: true,
              path: '/users/basics'
            },
            {
              title: 'Digital Wallets',
              directory: true,
              path: '/users/wallets'
            },
            {
              title: 'Account Keys',
              directory: true,
              path: '/users/keys'
            },
            {
              title: 'Technical Concepts',
              directory: true,
              path: '/users/technical_concepts'
            },
          ]
        },
        {
          title: 'For dApp Devs',
          children: [
            {
              title: 'Overview',
              directory: false,
              path: '/developers/overview'
            },
            {
              title: 'Quick Connect',
              directory: false,
              path: '/developers/connect'
            },
            {
              title: 'dapplink Clients',
              directory: false,
              path: '/developers/clients'
            },
            {
              title: 'Testnet Faucet',
              directory: false,
              path: '/developers/faucet'
            },
            {
              title: 'Ethereum Tooling',
              directory: true,
              path: '/developers/tools'
            },
            {
              title: 'Localnet',
              directory: true,
              path: '/developers/localnet'
            },
            {
              title: 'Libraries',
              directory: true,
              path: '/developers/libraries'
            },
            {
              title: 'Ethereum JSON-RPC',
              directory: true,
              path: '/developers/json-rpc'
            },
            {
              title: 'Cosmos gRPC & REST',
              path: 'https://api.dapplink.dev/'
            },
            {
              title: 'Tendermint RPC',
              path: 'https://docs.tendermint.com/v0.34/rpc/'
            },
          ]
        },
        {
          title: 'For Protocol Devs',
          children: [
            {
              title: 'Modules',
              directory: true,
              path: '/modules'
            },
            {
              title: 'Module Accounts',
              directory: false,
              path: '/protocol/moduleaccounts'
            },
            {
              title: 'Telemetry',
              directory: false,
              path: '/protocol/telemetry'
            },
            {
              title: 'IBC Channels',
              directory: false,
              path: '/protocol/ibc'
            }, 
            {
              title: 'dapplink Go API',
              path: 'https://pkg.go.dev/github.com/dapplink/dapplink'
            },
            {
              title: 'Ethermint Library Go API',
              path: 'https://pkg.go.dev/github.com/evmos/ethermint'
            },
            {
              title: 'dapplink Protobuf',
              directory: false,
              path: '/protocol/proto-docs'
            }
          ]
        },
        {
          title: 'For Validators',
          children: [
            {
              title: 'Validators Overview',
              directory: false,
              path: '/validators/overview'
            },
            {
              title: 'Setup & Configuration',
              directory: true,
              path: '/validators/setup'
            },
            {
              title: 'Installation & Quick Start',
              directory: true,
              path: '/validators/quickstart'
            },
            {
              title: 'Join Testnet',
              directory: false,
              path: '/validators/testnet'
            },
            {
              title: 'Join Mainnet',
              directory: false,
              path: '/validators/mainnet'
            },
            {
              title: 'Security',
              directory: true,
              path: '/validators/security'
            },
            {
              title: 'Software Upgrade Guide',
              directory: true,
              path: '/validators/upgrades'
            },
            {
              title: 'dapplink Governance',
              directory: true,
              path: '/validators/governance'
            },
            {
              title: 'Snapshots & Archive Nodes',
              directory: false,
              path: '/validators/snapshots_archives'
            },
            {
              title: 'FAQ',
              directory: false,
              path: '/validators/faq'
            }
          ]
        },
        {
          title: 'Block Explorers',
          children: [
            {
              title: 'Block Explorers',
              path: '/developers/explorers'
            },
            {
              title: 'Blockscout (EVM)',
              path: 'https://evm.dapplink.org'
            },
            {
              title: 'Mintscan (Cosmos)',
              path: 'https://www.mintscan.io/dapplink/'
            },
          ]
        },
      ]
    },
    gutter: {
      title: 'Help & Support',
      chat: {
        title: 'Discord Channel',
        text: 'Chat with dapplink users and team on Discord.',
        url: 'https://discord.gg/dapplink',
        bg: 'linear-gradient(103.75deg, #1B1E36 0%, #22253F 100%)'
      },
      forum: {
        title: 'Commonwealth Forum',
        text: 'Join the dapplink Commonwealth forum',
        url: 'https://commonwealth.im/dapplink',
        bg: 'linear-gradient(221.79deg, #3D6B99 -1.08%, #336699 95.88%)',
      },
      github: {
        title: 'Found an Issue?',
        text: 'Help us improve this page by suggesting edits on GitHub.',
        bg: '#F8F9FC'
      }
    },
    footer: {
      logo: '/dapplink-black.svg',
      textLink: {
        text: "dapplink.org",
        url: 'https://dapplink.org'
      },
      services: [
        {
          service: "github",
          url: 'https://github.com/dapplink/dapplink'
        },
        {
          service: "twitter",
          url: "https://twitter.com/dapplinkOrg",
        },
        {
          service: "telegram",
          url: "https://t.me/dapplinkOrg",
        },
        {
          service: "linkedin",
          url: "https://www.linkedin.com/company/tharsis-finance/",
        },
        {
          service: "medium",
          url: "https://dapplink.blog/",
        },
      ],
      smallprint: 'This website is maintained by Tharsis Labs Ltd.',
      links: [{
        title: 'Ecosystem Documentation',
        children: [
        {
          title: 'Cosmos SDK Docs',
          url: 'https://docs.cosmos.network'
        },
        {
          title: 'Ethereum Docs',
          url: 'https://ethereum.org/developers'
        },
        {
          title: 'Tendermint Core Docs',
          url: 'https://docs.tendermint.com'
        }
        ]
      },
      {
        title: 'Community',
        children: [
          {
            title: 'dapplink Discord Community',
            url: 'https://discord.gg/dapplink'
          },
          {
            title: 'dapplink Commonwealth Forum',
            url: 'https://commonwealth.im/dapplink'
          },
        ]
      },
      {
        title: 'dapplink',
        children: [
          {
            title: 'Jobs at dapplink',
            url: 'https://tharsis.notion.site/'
          }
        ]
      }
      ]
    },
    versions: [
      {
        "label": "main",
        "key": "main"
      },
    ],
  }
};
