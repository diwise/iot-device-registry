*** Settings ***
Documentation     A test suite that tests different Device creation scenarios.
Library           RequestsLibrary
Library           String
Resource          ../resources/fiware.robot

Suite Setup       suite setup


*** Test Cases ***
Create Device
    ${deviceID}=    Set Variable        urn:ngsi-ld:Device:${TEST_ID_PRFX}:mydevice
    ${device}=      Create Device       ${deviceID}  ${snowTempModel}
    ${resp}=        POST On Session     diwise      /ngsi-ld/v1/entities    json=${device}

    Status Should Be    201     ${resp}


Get Entities    
    ${resp}=        GET On Session      diwise      /ngsi-ld/v1/entities?type\=DeviceModel

    Status Should Be    200     ${resp}


Get Device
    ${deviceID}=    Set Variable        urn:ngsi-ld:Device:${TEST_ID_PRFX}:mydevice
    ${resp}=        GET On Session      diwise      /ngsi-ld/v1/entities/${deviceID}

    Status Should Be    200     ${resp}

    Entity Type And ID Should Match   ${resp.json()}  Device  ${deviceID}


Check That Update Entity Attributes Updates Values Correctly
    ${deviceID}=    Set Variable        urn:ngsi-ld:Device:${TEST_ID_PRFX}:mydevice
    ${resp}=            Update Device Value  diwise  ${deviceID}  t=10

    Update Device Value  diwise  ${deviceID}  snow=23
    Update Device Value  diwise  ${deviceID}  snow=24
    Update Device Value  diwise  ${deviceID}  t=11
    Update Device Value  diwise  ${deviceID}  t=12

    ${resp}=            Update Device Value  diwise  ${deviceID}  snow=25
    Status Should Be    204     ${resp}
    
    ${resp}=            GET On Session          diwise      /ngsi-ld/v1/entities/${deviceID}
    ${deviceValue}=     Get From Dictionary     ${resp.json()}     value
    ${value}=           Get From Dictionary     ${deviceValue}     value

    Should Be Equal As Strings   snow%3D25%3Bt%3D12  ${value}

Check That Update Entity Attributes Updates Location Correctly
    ${deviceID}=    Set Variable        urn:ngsi-ld:Device:${TEST_ID_PRFX}:mydevice
    ${longitude}=       Convert To Number        0.0 
    ${latitude}=        Convert To Number        0.0
    ${resp}=            Update Device Location  diwise  ${deviceID}  ${latitude}  ${longitude}

    Update Device Location      diwise  ${deviceID}  ${latitude}  ${longitude}

    ${longitude}=       Convert To Number        17.3069 
    ${latitude}=        Convert To Number        62.3908

    ${resp}=            Update Device Location  diwise  ${deviceID}  ${latitude}  ${longitude}    
    Status Should Be    204     ${resp}
    
    ${resp}=            GET On Session          diwise      /ngsi-ld/v1/entities/${deviceID}
    ${deviceLocation}=     Get From Dictionary     ${resp.json()}     location
    ${location}=           Get From Dictionary     ${deviceLocation}     value
    ${coordinates}=        Get From Dictionary     ${location}     coordinates

    Should Be Equal As Strings  [17.3069, 62.3908]   ${coordinates}


Check That On Off Values Can Be Handled
    ${deviceID}=    Set Variable        urn:ngsi-ld:Device:${TEST_ID_PRFX}:livboj
    ${device}=      Create Device       ${deviceID}  ${onOffModel}
    ${resp}=        POST On Session     diwise      /ngsi-ld/v1/entities    json=${device}

    Status Should Be    201     ${resp}
    ${resp}=        Update Device Value  diwise  ${deviceID}  on

    Status Should Be    204     ${resp}
    
    ${resp}=            GET On Session          diwise      /ngsi-ld/v1/entities/${deviceID}

    ${deviceValue}=     Get From Dictionary     ${resp.json()}     value
    ${value}=           Get From Dictionary     ${deviceValue}     value

    Should Be Equal As Strings   ${value}     on


Check Date Last Value Reported Updates Correctly
    ${deviceID}=    Set Variable        urn:ngsi-ld:Device:${TEST_ID_PRFX}:mydevice

    ${resp}=            Update Device Value  diwise  ${deviceID}  t=10
    ${resp}=            GET On Session       diwise  /ngsi-ld/v1/entities/${deviceID}
    ${firstDateValueReported}=            Get From Dictionary  ${resp.json()}       dateLastValueReported

    #temporary fix:
    BuiltIn.Sleep       1s

    ${resp}=            Update Device Value  diwise  ${deviceID}  t=25
    ${resp}=            GET On Session       diwise  /ngsi-ld/v1/entities/${deviceID}
    ${lastDateValueReported}=           Get From Dictionary     ${resp.json()}      dateLastValueReported

    Should Not Be Equal As Strings      ${firstDateValueReported}      ${lastDateValueReported}


Get Device Model
    ${deviceModelID}=  Set Variable     urn:ngsi-ld:DeviceModel:${TEST_ID_PRFX}:snowtemp
    ${resp}=           GET On Session      diwise    /ngsi-ld/v1/entities/${deviceModelID}
    Status Should Be   200     ${resp}

    Entity Type And ID Should Match  ${resp.json()}  DeviceModel  ${deviceModelID}


*** Keywords ***
suite setup
    ${headers}=       Create Dictionary   Content-Type=application/ld+json
    Create Session    diwise    http://127.0.0.1:8686  headers=${headers}

    ${TEST_ID_PRFX}=  Generate Random String  8  [NUMBERS]abcdef
    Set Suite Variable  ${TEST_ID_PRFX}

    ${modelID}=         Set Variable        urn:ngsi-ld:DeviceModel:${TEST_ID_PRFX}:snowtemp
    ${snowTempModel}=   Create Device Model For Properties  ${modelID}  temperature  snowDepth
    ${resp}=            POST On Session     diwise      /ngsi-ld/v1/entities    json=${snowTempModel}
    Status Should Be    201  ${resp}
    Set Suite Variable  ${snowTempModel}

    ${modelID}=         Set Variable        urn:ngsi-ld:DeviceModel:${TEST_ID_PRFX}:onoff
    ${onOffModel}=      Create Device Model For Properties  ${modelID}  state
    ${resp}=            POST On Session     diwise      /ngsi-ld/v1/entities    json=${onOffModel}
    Status Should Be    201  ${resp}
    Set Suite Variable  ${onOffModel}
