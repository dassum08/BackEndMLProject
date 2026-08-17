from typing import Generator
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, Session
from sqlalchemy.ext.declarative import declarative_base
import os

DATABASE_URL = os.getenv("DATABASE_URL")
#DATABASE_URL =  "postgresql://postgres:root@localhost:5432/MLDB"

Base = declarative_base()
# Create engine with connection pooling parameters optimized for PostgreSQL
engine = create_engine(
    DATABASE_URL,
    pool_pre_ping=True,  # Automatically tests/recycles stale database connections
    pool_size=5,         # Base connection pool limit
    max_overflow=10      # Max burst connections beyond pool_size
)

# Session factory for generating clean stateful transactions
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

def get_db() -> Generator[Session, None, None]:
    """
    FastAPI Dependency that yields a clean database session context per request.
    Guarantees session closure even if errors crop up during route execution.
    """
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
