from fastapi import APIRouter, Depends
from app.schema import PredictionRequest, PredictionResponse, PredictionModel
from app.model import predict, FireCompPred
from sqlalchemy.orm import Session
from app.db.session import get_db, engine, Base
 
router = APIRouter()

Base.metadata.create_all(bind=engine)

@router.post("/predict", response_model=PredictionResponse)
def make_prediction(request: PredictionRequest, db: Session = Depends(get_db)):
    features = [
        request.smoke_detector,
        request.new_batteries,
        request.abc_extinguisher,
        request.clear_exit_routes
    ]
    print(request.smoke_detector,request.new_batteries,request.abc_extinguisher,request.clear_exit_routes)
    class_id, probability = predict(features)
    

    db_fire = FireCompPred(
        smoke_detector = request.smoke_detector,
        new_batteries = request.new_batteries,
        abc_extinguisher = request.abc_extinguisher,
        clear_exit_routes = request.clear_exit_routes,
        prediction = class_id,
        probabilities = probability 
    )
    
    db.add(db_fire)
    db.commit()
    db.refresh(db_fire)  # Refreshes instance attributes with newly assigned database fields (like id)
    
    print("class_id:",class_id)
    print("probability:",probability)
    return PredictionResponse(
        class_id=class_id,
        probability=probability
      )


@router.post("/getdata", response_model=PredictionModel)
def postdata(db: Session = Depends(get_db)):
    return db.query(FireCompPred).all()[-1]      