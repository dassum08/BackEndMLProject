from pydantic import BaseModel, ConfigDict

class PredictionRequest(BaseModel):
    smoke_detector: int
    new_batteries : int
    abc_extinguisher: int
    clear_exit_routes: int

class PredictionResponse(BaseModel):
    class_id: int
    probability: float
    
class PredictionModel(BaseModel):
    smoke_detector: int
    new_batteries : int
    abc_extinguisher: int
    clear_exit_routes: int
    prediction: int
    probabilities: float
    
    model_config = ConfigDict(from_attributes=True)