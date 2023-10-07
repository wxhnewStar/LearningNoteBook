import pandas as pd

# 创建数据帧
data = {'Name': ['Alice', 'Bob', 'Charlie'],
        'Age': [25, 30, 35],
        'City': ['New York', 'London', 'Paris']}
df = pd.DataFrame(data)

# 查看数据
print(df.head())

# 访问数据
print(df['Name'])
print(df.loc[1])

# 添加新列
df['Salary'] = [50000, 60000, 70000]

# 删除列
df.drop('City', axis=1, inplace=True)

# 修改数据
df.loc[0, 'Age'] = 26

# 过滤数据
filtered_df = df[df['Age'] > 30]

# 排序数据
sorted_df = df.sort_values(by='Age')

# 聚合数据
avg_age = df['Age'].mean()

# 显示数据帧
print(df)