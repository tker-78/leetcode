# leetcode

LeetCodeの学習記録用レポジトリです。  
GolangかPython3で解答しています。  
たまにSQLの問題も解答するかもしれません。  


## Pandas todo

- lambdaの使い方を確認する

```python
employees["bonus"] = employees.apply(
    lambda x: x["salary"] * 2 if x["employee_id"] % 2 else 0,
    axis=1
)
```

