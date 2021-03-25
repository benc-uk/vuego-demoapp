module.exports = {
  env: {
    node: true,
    es6: true
  },
  extends: ['eslint:recommended', 'plugin:vue/recommended'],
  ignorePatterns: ['gauge.min.js'],
  rules: {
    // Errors & best practices
    'no-var': 'error',
    'no-console': 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-unused-vars': ['error', { argsIgnorePattern: 'next|res|req' }],
    curly: 'error',

    // ES6
    'arrow-spacing': 'error',
    'arrow-parens': 'error',

    // Vue
    'vue/html-closing-bracket-newline': 'off',
    'vue/max-attributes-per-line': 'off',
    'vue/html-self-closing': 'off',
    'vue/singleline-html-element-content-newline': 'off'
  },
  parserOptions: {
    parser: 'babel-eslint'
  }
}
