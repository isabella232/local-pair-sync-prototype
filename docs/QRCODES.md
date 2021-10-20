# QR Codes

**Terms**:
- Sender = The device sending private key data
- Receiver = The device receiving private key data

## Scenarios

Sharing:

- AES key
- Device UUID
- x509 cert

### Share QR code from Sender

**Before the scan :**

The Receiver knows:

- It is connected with a device that claims to be the Sender
- The Sender device has presented a x509 cert

The Receiver DOES NOT know:

- The Sender x509 cert is signed by a trusted key
- That the Sender is a valid device

The Sender knows:

- It is connected with a device that claims to be the Receiver

The Sender DOES NOT know:

- That the Receiver is a valid device

**After the QR scan :**

The Receiver knows:

- Whether the x509 cert's public key matches what the Sender has issued
  - If public key matches we can consider the x509 cert as authorised
  - If public key does not match we have a potential MITM attack
- The encryption key for the private key payload

The Sender knows:

- It is connected with a device that claims to be the Receiver
- That it has an encryption key shared only with the Receiver
  - Therefore, only the Receiver can decrypt the private key payload

The Sender DOES NOT know:

- That the Receiver is a valid device
  - But only the Receiver can decrypt the private key payload
  
### Share QR code from Receiver

**Before the scan :**

The Receiver knows:

- It is connected with a device that claims to be the Sender
- The Sender device has presented a x509 cert

The Receiver DOES NOT know:

- The Sender x509 cert is signed by a trusted key
- That the Sender is a valid device

The Sender knows:

- It is connected with a device that claims to be the Receiver

The Sender DOES NOT know:

- That the Receiver is a valid device

**After the QR scan :**

The Receiver knows:

- It is connected with a device that claims to be the Sender
- The Sender device has presented a x509 cert
- That it has an encryption key shared only with the Sender
  - Therefore, any payload coming from the device claiming to be the Sender can only be decrypted by the shared AES key
  - This means that if the payload from the device claiming to be the Sender can not be decrypted by the shared key the payload and connection must be discarded

The Receiver **DOES NOT** know:

- The Sender x509 cert is signed by a trusted key
- That the Sender is a valid device

The Sender knows:

- It is connected with a device that claims to be the Receiver
- Whether the Receiver has access to the valid Sender x509 cert
  - If the public key from the QR code matches the ephemeral key the Sender knows the Receiver was able to access its x509 cert
  - If the public key from the QR code does not match, the Sender must discard the connection with the Receiver
- That it has an encryption key shared only with the Receiver
    - Therefore, only the Receiver can decrypt the private key payload

The Sender DOES NOT know:

- That the Receiver is a valid device
    - But only the Receiver can decrypt the private key payload
 
