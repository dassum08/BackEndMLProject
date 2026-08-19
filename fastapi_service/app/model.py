import joblib
from sqlalchemy import Column, Integer, Float
from app.db.session import Base

MODEL_PATH = "app/model.joblib"
model = joblib.load(MODEL_PATH)

def predict(features):
    prediction = model.predict([features])[0]
    probabilities = model.predict_proba([features])[0]
    return int(prediction), float(max(probabilities))
    
 

class FireCompPred(Base):
    __tablename__ = "fire_compliance"
    id = Column(Integer, primary_key=True, index=True)
    smoke_detector = Column(Integer, index=True, nullable=False)
    new_batteries = Column(Integer, index=True, nullable=False)
    abc_extinguisher = Column(Integer, index=True, nullable=False)
    clear_exit_routes = Column(Integer, index=True, nullable=False)
    prediction = Column(Integer, index=True, nullable=False)
    probabilities = Column(Float, index=True, nullable=False)