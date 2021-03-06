/*
Copyright DTCC 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

         http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package org.hyperledger.java.shim;

import com.google.protobuf.ByteString;
import io.grpc.stub.StreamObserver;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.hyperledger.java.fsm.CBDesc;
import org.hyperledger.java.fsm.Event;
import org.hyperledger.java.fsm.EventDesc;
import org.hyperledger.java.fsm.FSM;
import org.hyperledger.java.fsm.exceptions.CancelledException;
import org.hyperledger.java.fsm.exceptions.NoTransitionException;
import org.hyperledger.java.helper.Channel;
import org.hyperledger.protos.Chaincode.*;
import org.hyperledger.protos.Chaincodeshim.*;
import org.hyperledger.protos.Chaincodeshim.ChaincodeMessage.Builder;
import org.hyperledger.protos.ProposalResponseOuterClass.Response;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static org.hyperledger.java.fsm.CallbackType.*;
import static org.hyperledger.protos.Chaincodeshim.ChaincodeMessage.Type.*;


public class ChaincodeResponse {

	private static Log logger = LogFactory.getLog(ChaincodeResponse.class);
	
    private static int OK = 200; 

    private static int ERROR = 500; 

	public ChaincodeResponse() {
		
	}
    
    public static ByteString Success(String payload){
        logger.info("payload:"+payload);
        Response res = Response
        .newBuilder()
        .setStatus(OK)
        .setPayload(ByteString.copyFromUtf8(payload))
        .build();
        //.toByteString()
        //.toStringUtf8(); 
        //logger.info("Success:"+res);
        return res.toByteString();

    }

    public static ByteString Error(String msg){
        logger.info("msg:"+msg);
        Response res = Response
        .newBuilder()
        .setStatus(ERROR)
        .setMessage(msg)
        .build();
        //.toByteString()
        //.toStringUtf8();
        //logger.info("Error:"+res);
        return res.toByteString();
    }
	
}
