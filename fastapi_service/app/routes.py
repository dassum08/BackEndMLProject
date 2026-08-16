from fastapi import APIRouter, Depends
from app.schema import PredictionRequest, PredictionResponse
from app.model import predict, IrisPred
from sqlalchemy.orm import Session
from app.db.session import get_db, engine, Base
router = APIRouter()

Base.metadata.create_all(bind=engine)

@router.post("/predict", response_model=PredictionResponse)
def make_prediction(request: PredictionRequest, db: Session = Depends(get_db)):
    features = [
        request.sepal_length,
        request.sepal_width,
        request.petal_length,
        request.petal_width
    ]
    class_id, probability = predict(features)
    

    db_iris = IrisPred(
        sepal_length = request.sepal_length,
        sepal_width = request.sepal_width,
        petal_length = request.petal_length,
        petal_width = request.petal_width,
        prediction = class_id,
        probabilities = probability 
    )
    
    db.add(db_iris)
    db.commit()
    db.refresh(db_iris)  # Refreshes instance attributes with newly assigned database fields (like id)
    
    return PredictionResponse(
        class_id=class_id,
        probability=probability
      )  