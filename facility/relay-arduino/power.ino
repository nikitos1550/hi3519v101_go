
int latchPin = 2;
int clockPin = 3;
int dataPin = 4;

void setup() {
  pinMode(latchPin, OUTPUT);
  pinMode(clockPin, OUTPUT);
  pinMode(dataPin, OUTPUT);

  sendout();

  Serial.begin(115200); 
  Serial.println("READY");
}

unsigned int output = 0xFFFF;//default all on

void sendout() {
    unsigned int send = output;
    digitalWrite(latchPin, LOW);
    shiftOut(dataPin, clockPin, MSBFIRST, send >> 8);
    shiftOut(dataPin, clockPin, MSBFIRST, (send << 8) >> 8);
    digitalWrite(latchPin, HIGH);
}

void relayon(unsigned int num) {
    unsigned int tmp = (1 << (num - 1));
    output = output | tmp;
    sendout();
}

void relayoff(unsigned int num) {
    unsigned int tmp = (1 << (num - 1));
    output = output & ~tmp;
    Serial.println(output, BIN);
    sendout();
}

unsigned char count = 0;
char buf[255];

void loop() {
  if (Serial.available() > 0) {  //если есть доступные данные
        //Serial.println(Serial.available());
        
        buf[count] = Serial.read();
        //Serial.println(buf[count], HEX);
        
        if (buf[count] == '\n') {
          //Serial.println("got");
          
          buf[count] = '\0';
          //Serial.println(buf);
          
          if        (memcmp (buf, "on",     2) == 0) {
            Serial.println("ON");
            unsigned char num = buf[2+1] - '0';
            Serial.println(num);
            relayon(num);
          } else if (memcmp (buf, "off",    3) == 0) {
            Serial.println("OFF");
            unsigned char num = buf[3+1] - '0';
            relayoff(num);
            Serial.println(num);
          } else if (memcmp (buf, "reset",  5) == 0) {
            Serial.println("RESET");
            unsigned char num = buf[5+1] - '0';
            relayoff(num);
            delay(1000);
            relayon(num);
            Serial.println(num);            
          } else {
            Serial.println("Unrecognized command");
          }
          count=0;
        } else {
          count++;
        }
  }
}
