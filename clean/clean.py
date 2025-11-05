import pandas as pd
import os

INPUT_FILE = os.getenv("INPUT_FILE")
OUTPUT_FILE = os.getenv("INPUT_FILE")

# === 1. Load the Bitso CSV ===
# Skip the first messy row containing the "Created on" and URL junk
bitso = pd.read_csv(INPUT_FILE, skiprows=1)

# === 2. Clean column names ===
bitso.columns = [c.strip().lower().replace(' ', '_') for c in bitso.columns]

# === 3. Keep only the important columns ===
# (Drop extra ones like symbol, volume_mxn, etc.)
keep_cols = ['timestamp', 'open', 'high', 'low', 'close', 'volume_btc']
bitso = bitso[keep_cols]

# === 4. Rename columns to match BTCUSDT_1m_large.csv ===
bitso = bitso.rename(columns={'volume_btc': 'volume'})

# === 5. (Optional) Sort by timestamp if needed ===
bitso = bitso.sort_values(by='timestamp')

# === 6. Save the cleaned version ===
bitso.to_csv(OUTPUT_FILE, index=False)

print("✅ Cleaned CSV saved as Bitso_BTCMXN_d_cleaned.csv")
print(bitso.head())
