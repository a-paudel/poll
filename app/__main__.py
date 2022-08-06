import dotenv

dotenv.load_dotenv()

import os
from app.app import app

app.run(
    debug=bool(os.getenv("DEBUG", False)),
    port=os.getenv("PORT", 5000),
    host=os.getenv("HOST", "0.0.0.0"),
)
