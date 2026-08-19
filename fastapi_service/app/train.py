from sklearn.datasets import load_iris
from sklearn.ensemble import RandomForestClassifier
import joblib
import pandas as pd

# Load dataset
#iris = load_iris()

#X = iris.data
#y = iris.target

df = pd.read_csv('fire_data.csv')

# 2. Split into X (all columns except the last) and y (only the last column)
X = df.iloc[:, :-3]
y = df.iloc[:, -1]

# Train model
model = RandomForestClassifier(
    n_estimators=100,
    random_state=42
)

model.fit(X, y)

# Save model
joblib.dump(model, "model.joblib")

print("Model trained and saved.")