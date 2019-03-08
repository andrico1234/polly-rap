#!/bin/bash

if [ -f ../out.test.mp3 ]; then
  rm ../out/test.mp3
fi

aws polly synthesize-speech \
  --output-format mp3 \
  --text "I'm a big four-eyed lame-o and I wear the same stupid sweater everyday, and - THE SPRINGFIELD RIVER!" \
  --voice-id Kendra \
  ../out/test.mp3
