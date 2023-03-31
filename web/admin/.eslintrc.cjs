module.exports = {
  root: true,
  env: {
    node: true,
  },
  parser: '@typescript-eslint/parser',
  plugins: ['react', '@typescript-eslint'],
  rules: {
    // 强制在注释中 // 或 /* 使用一致的空格 http://eslint.cn/docs/rules/spaced-comment
    'spaced-comment': [2],
    // 强制在 function的左括号之前使用一致的空格  http://eslint.cn/docs/rules/space-before-function-paren
    // 'space-before-function-paren': [2, 'always'],
    // 禁止不规则的空白/允许在模板字面量中出现任何空白字符  http://eslint.cn/docs/rules/no-irregular-whitespace
    'no-irregular-whitespace': [2, {'skipTemplates': true}],
    // switch语句必须包含default http://eslint.cn/docs/rules/default-case
    'default-case': 'error',
    // 禁止出现多个空格 http://eslint.cn/docs/rules/no-multi-spaces
    'no-multi-spaces': 'error',
    // 禁止自身比较 http://eslint.cn/docs/rules/no-self-compare
    'no-self-compare': 'error',
    // 禁止多余的return http://eslint.cn/docs/rules/no-useless-return
    'no-useless-return': 'error',
    // 禁止使用不带 await 表达式的 async 函数 http://eslint.cn/docs/rules/require-await
    'require-await': 'error',
    // 数组括号前后不允许出现空格 http://eslint.cn/docs/rules/array-bracket-spacing
    'array-bracket-spacing': [2, 'never'],
    // 大括号 强制 one true brace style/允许开括号和闭括号在同一行 http://eslint.cn/docs/rules/brace-style
    'brace-style': [2, '1tbs', {'allowSingleLine': true}],
    // 强制逗号前后使用一致的空格 http://eslint.cn/docs/rules/comma-spacing
    'comma-spacing': [2],
    // 强制在计算的属性的方括号中使用一致的空格(不使用) http://eslint.cn/docs/rules/computed-property-spacing
    'computed-property-spacing': [2, 'never'],
    // 禁止在函数标识符和其调用之间有空格 http://eslint.cn/docs/rules/func-call-spacing
    'func-call-spacing': [2, 'never'],
    // 强制在对象字面量的属性中键和值之间使用一致的间距 http://eslint.cn/docs/rules/key-spacing
    'key-spacing': [2],
    // 强制在关键字前后使用一致的空格 http://eslint.cn/docs/rules/keyword-spacing
    'keyword-spacing': [2, { 'before': true, 'after': true }],
    // 强制在块前使用一致的空格 http://eslint.cn/docs/rules/space-before-blocks
    'space-before-blocks': [2, 'always'],
    // 要求操作符前后有空格 http://eslint.cn/docs/rules/space-infix-ops
    'space-infix-ops': [2],

    quotes: [1, 'single'], // 使用单引号
    semi: [2, 'never'], // 结尾不使用分号
    'prefer-const': 0, // 首选const
    'indent': ['error', 2, { 'ignoredNodes': ['VariableDeclaration[declarations.length=0]'] }],
    // "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-unused-vars': 0,
    'jsx-quotes': ['error', 'prefer-double'],
    'react/self-closing-comp': ['error']
  }
}