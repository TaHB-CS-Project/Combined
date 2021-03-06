
$.getJSON("js/record-list.json", 
    function (data) {
    var info = '';
    var modal_moreinfo = '';
    var modal_editinfo = '';

    $.each(data, function (key, value) {
        var moreinfo = '';
        var editinfo = '';
        var modalcreate_set_name = 'id="c' + value.Record_id + '"';
        var modalcreate_name = '#c' + value.Record_id;
        var modalbody_set_name = 'id="d' + value.Record_id + '"';
        var modalbody_name = '#d' + value.Record_id;
        var modalcreate_set_name_edit = 'id="e' + value.Record_id + '"';
        var modalcreate_name_edit = '#e' + value.Record_id;
        var modalbody_set_name_edit = 'id="t' + value.Record_id + '"';
        var modalbody_name_edit = '#t' + value.Record_id;
        var hospital = 'id="hospital' + value.Record_id + '"';
        var hospital_name = '#hospital' + value.Record_id; 
 //     var hospital_id = '"hospital' + value.Record_id +'"';
        var diagnosis = 'id="diagnosis' + value.Record_id + '"';
        var diagnosis_name = '#diagnosis' + value.Record_id;
 //      var diagnosis_id = '"diagnosis' + value.Record_id + '"';
        var procedure = 'id="procedure' + value.Record_id + '"';
        var procedure_name = '#procedure' + value.Record_id;
//      var procedure_id = '"procedure' + value.Record_id + '"';
        var gender = 'id="gender' + value.Record_id + '"';
        var record_birthday = 'id="record_birthday' + value.Record_id + '"';
        var record_date = 'id="record_date' + value.Record_id + '"';
        var record_id = 'id="record_id' + value.Record_id + '"';
        var record_id_name = '#record_id' +value.Record_id;
        var result = 'id="result' + value.Record_id + '"';
        var special_notes = 'id="special_notes' + value.Record_id + '"';
        var weight = 'id="weight' + value.Record_id + '"';
        

        info += '<tr>';
        info += '<td>' + value.Record_id + '</td>';
        info += '<td id=' + value.Hospital_name + '>' + value.Hospital_name + '</td>';
        info += '<td id=' + value.Start_datetime + '>' + value.Start_datetime + '</td>';
/*        info += '<td id=' + value.Medicalemployee_firstname + '>' + 
            value.Medicalemployee_firstname + '</td>';
*/           
        info += '<td id=' + value.Medicalemployee_lastname + '>' + value.Medicalemployee_lastname + '</td>';
        info += '<td id=' + value.Procedure_name + '>' + value.Procedure_name + '</td>';          
 /*       info += '<td id=' + value.Diagnosis_name + '>' + 
            value.Diagnosis_name + '</td>';
 */       
        info += '<td id=' + value.Outcome + '>' + value.Outcome + '</td>';
        info += `<td>
                <button type="button" class="btn btn-primary" data-toggle="modal" data-target="`+ modalcreate_name +`">Details</button>
                <button type="button" class="btn btn-primary" data-toggle="modal" data-target="`+ modalcreate_name_edit +`">Edit</button>
                </td>`;
        info += '</tr>'

        modal_moreinfo += `<div class="modal fade" ` + modalcreate_set_name + `=" tabindex="-1" role="dialog" aria-labelledby="exampleModalLongTitle" aria-hidden="true">
                <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                    <h5 class="modal-title" id="exampleModalLongTitle">Record `+ value.Record_id +`</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                    </div>
                    <div class="modal-body text-center" `+ modalbody_set_name +`">

                    </div>
                    <div class="modal-footer">
                    <button type="button" id="button`+ value.Record_id +`" class="btn btn-danger">Delete</button>
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                    </div>
                    </div>
                </div>
                </div>
                </div>`;

        moreinfo += '<div class="flex-row">'
        moreinfo += `<div class="form-group row">
                        <label for="colFormLabel" class="col-sm-6 col-form-label"><b>Record ID:</b></label>
                        <div class="col-sm-6" >
                        <input style="text-align:center" type="text" class="form-control" `+ record_id +` name="record_id" value="`+ value.Record_id +`" readonly>
                        </div></div>`
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Hospital: </b></div><div class="col-sm-6">'+ value.Hospital_name+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Date: </b></div><div class="col-sm-6">'+ value.Start_datetime+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Dortor Firstname: </b></div><div class="col-sm-6">'+ value.Medicalemployee_firstname+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Doctor Lastname: </b></div><div class="col-sm-6">'+ value.Medicalemployee_lastname+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Procedure: </b></div><div class="col-sm-6">'+ value.Procedure_name+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Diagnosis: </b></div><div class="col-sm-6">'+ value.Diagnosis_name+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Outcome: </b></div><div class="col-sm-6">'+ value.Outcome+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Special Notes: </b></div><div class="col-sm-6">'+ value.Special_notes+'</div></div>'
        moreinfo += '<br><b>============== Patient Information ==============</b><br><br>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Patient Birthday: </b></div><div class="col-sm-6">'+ value.Patient_birthday+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b>  Patient Sex: </b></div><div class="col-sm-6">'+ value.Patient_sex+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Patient Weight: </b></div><div class="col-sm-6">'+ value.Patient_weightlbs+' lbs</div></div>'
        moreinfo += '</div>'


        modal_editinfo += `<div class="modal fade" ` + modalcreate_set_name_edit + `=" tabindex="-1" role="dialog" aria-labelledby="exampleModalLongTitle" aria-hidden="true">
                <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                    <h5 class="modal-title" id="exampleModalLongTitle">Record `+ value.Record_id +`</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                    </div>
                    <div class="modal-body text-center" `+ modalbody_set_name_edit +`">
                    </div>
                </div>
                </div>
                </div>`;

        editinfo += `<form action="/submit_record_draft" method="POST">
                            <div class="form-group row">
                            <label for="colFormLabel" class="col-sm-4 col-form-label">Record ID</label>
                            <div class="col-sm-8" >
                            <p class="form-control" `+ record_id +`>`+ value.Record_id +` </p>
                            </div>
                        </div>
                        <div class="form-group row">
                            <label for="colFormLabel" class="col-sm-4 col-form-label">Hospital</label>
                            <div class="col-sm-8">
                            <select type="name" class="form-control" `+ hospital +`name="hospital">
                            <option value="`+ value.Hospital_name +`">`+ value.Hospital_name +`</option>
                            <option>----------</option>
                            </select>
                            </div>
                        </div>
                        <div class="form-group row">
                            <label for="colFormLabel" class="col-sm-4 col-form-label">Date</label>
                            <div class="col-sm-8" >
                                <input type="text" class="form-control" data-date-format="yyyy-mm-dd" name="record_date" `+ record_date +` value="`+ value.Start_datetime +`" placeholder="yyyy-mm-dd">
                            </div>
                        </div>
                        <div class="form-group row">
                            <label for="colFormLabel" class="col-sm-4 col-form-label">Gender</label>
                            <div class="col-sm-8">
                            <select type="name" class="form-control" `+ gender +` name="gender">
                                <option value="`+ value.Patient_sex +`">`+ value.Patient_sex +`</option>
                                <option>-----------</option>
                                <option value="Male">Male</option>
                                <option value="Female">Female</option>
                                </select>
                            </div>
                        </div>
                        <div class="form-group row">
                            <label for="colFormLabel" class="col-sm-4 col-form-label">Weight</label>
                                <div class="col-sm-7">
                                    <input type="name" class="form-control" `+ weight +` placeholder="Weight" name="weight" value="`+ value.Patient_weightlbs +`">
                                </div>
                                <div class="col-sm-1 align-self-center">
                                    <span>lb</span>
                                </div>
                        </div>
                        
                        <div class="form-group row">
                            <label for="colFormLabel" class="col-sm-4 col-form-label">Date of Birth</label>
                            <div class="col-sm-8" >
                            <input type="text" class="form-control"  name="record_birthday" `+ record_birthday +` value="`+ value.Patient_birthday +`"placeholder="yyyy-mm-dd">
                        </div>
                        </div>
                    
                        <div class="form-group row">
                            <label for="colFormLabel" class="col-sm-4 col-form-label">Diagnosis</label>
                            <div class="col-sm-8">
                                <select name="diagnosis" class="form-control" `+ diagnosis +`>
                                <option value="`+ value.Diagnosis_name +`">`+ value.Diagnosis_name +`</option>
                                <option>-----------</option>
                                </select>
                            </div>
                        </div>
                    
                        <div class="form-group row">
                        <label for="colFormLabel" class="col-sm-4 col-form-label">Procedure</label>
                        <div class="col-sm-8 procedure_wrapper">
                            <select name="procedure" class="form-control" `+ procedure +`>
                            <option value="`+ value.Procedure_name +`">`+ value.Procedure_name +`</option>
                            <option>-----------</option>
                    
                            </select>
                        </div>
                        </div>
                    
                    
                    <div class="form-group row">
                        <label for="colFormLabel" class="col-sm-4 col-form-label">Result/Outcome</label>
                        <div class="col-sm-8">
                        <select type="name" class="form-control" `+ result +` name="result">
                            <option value="`+ value.Outcome +`">`+ value.Outcome +`</option>
                            <option>-----------</option>
                            <option value="Success">Success</option>
                            <option value="Complication">Complication</option>
                            <option value="Inconclusive">Inconclusive</option>
                            </select>
                        </div>
                    </div>
                    <div class="form-group row">
                        <label for="colFormLabel" class="col-sm-4 col-form-label">Special Notes</label>
                        <div class="col-sm-8">
                        <textarea class="form-control" `+ special_notes +` rows="4" cols="50" placeholder="Special Notes" name="special_notes">`+ value.Special_notes+`</textarea> 
                        </textarea>
                        </div>
                    </div>
                    <div class="modal-footer">
                    <button type="submit" id="submit_draft_delete" class="btn btn-primary">Submit</button>
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                    </div>
                    </div>
                    </form>`     

    $.getJSON("js/hospitalnamelist.json",
     function (data) {
         $.each(data,
         function (key, value) {
         $(hospital_name).append("<option value='" + value.Hospital_name + "'>" + value.Hospital_name + "</option>");
         });
     });
     
    $.getJSON("js/diagnosis.json",
     function (data) {
         $.each(data,
         function (key, value) {
         $(diagnosis_name).append("<option value='" + value.Diagnosis_name + "'>" + value.Diagnosis_name + "</option>");
         });
     });
     
    $.getJSON("js/procedure.json",
     function (data) {
         $.each(data,
         function (key, value) {
         $(procedure_name).append("<option value='" + value.Procedure_name + "'>" + value.Procedure_name + "</option>");
            });
    });

    $('#submit_draft_delete').submit(function(event){
        event.preventDefault();
        var record=$(record_id_name).val();

            $.ajax(  
                {
                    url:'/submit_draft_delete',    
                    type:"POST",   
                    dataType:"JSON", 
                    data: JSON.stringify({"Record_draft_id": record}),
                    success:function(){  
                        alert("Succesfully Submitted!")
                    },
                    error: function() { 
                        alert("Somethings Wrong"); 
                    }    

                }
            );
        });
    
        $('#modal_moreinfo').append(modal_moreinfo);
        $('#modal_editinfo').append(modal_editinfo);
        $(modalbody_name).append(moreinfo);
        $(modalbody_name_edit).append(editinfo);
       
     });

     $('#record_list_datatable').append(info);
     
});

