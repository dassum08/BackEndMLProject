import joblib
from sqlalchemy import Column, Integer, Float
from app.db.session import Base

MODEL_PATH = "app/model.joblib"
model = joblib.load(MODEL_PATH)

def predict(features):
    prediction = model.predict([features])[0]
    probabilities = model.predict_proba([features])[0]
    return int(prediction), float(max(probabilities))
    
 

class IrisPred(Base):
    __tablename__ = "iris_pred"
    id = Column(Integer, primary_key=True, index=True)
    sepal_length = Column(Float, index=True, nullable=False)
    sepal_width = Column(Float, index=True, nullable=False)
    petal_length = Column(Float, index=True, nullable=False)
    petal_width = Column(Float, index=True, nullable=False)
    prediction = Column(Integer, index=True, nullable=False)
    probabilities = Column(Float, index=True, nullable=False)