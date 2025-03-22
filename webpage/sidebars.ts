import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';
import apiSidebar from './docs/api/sidebar';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */
const sidebars: SidebarsConfig = {
  docsSidebar: [
    {
      type: 'category',
      label: 'Getting Started',
      className: 'menu__list-item-flat',
      collapsible: false,
      items: [{ type: "autogenerated", dirName: "getting-started" }]
    },
    {
      type: 'category',
      label: 'Reference',
      className: 'menu__list-item-flat',
      collapsible: false,
      items: [
        { type: "autogenerated", dirName: "reference" },
        {
          type: 'link',
          label: 'API Reference',
          href: '/docs/v3/api',
        }
      ]
    },
    //{
    //  type: 'category',
    //  label: 'Hardware Acceleration',
    //  items: [{ type: "autogenerated", dirName: "hardware-acceleration" }]
    //},
    //{
    //  type: 'category',
    //  label: 'Application Customization',
    //  items: [{ type: "autogenerated", dirName: "app-customization" }]
    //},
    //{
    //  type: 'category',
    //  label: 'Developer Guide',
    //  items: [{ type: "autogenerated", dirName: "developer-guide" }]
    //},
    {
      type: 'category',
      label: 'Help & Support',
      className: 'menu__list-item-flat',
      collapsible: false,
      items: [
        'faq',
        'release-notes',
        'roadmap',
      ]
    },
  ],
  apiSidebar: require("./docs/api/sidebar.js"),
};

export default sidebars;
