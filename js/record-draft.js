
$.getJSON("js/record-draft-list.json", 
function (data) {
var info = '';
var modal_editinfo = '';

$.each(data, function (key, value) {

    var editinfo = '';
 
    var modalcreate_set_name_edit = 'id="e' + value.Record_id + '"';
    var modalcreate_name_edit = '#e' + value.Record_id;
    var modalbody_set_name_edit = 'id="t' + value.Record_id + '"';
    var modalbody_name_edit = '#t' + value.Record_id;
    var hospital = 'id="hospital' + value.Record_id + '"';
    var hospital_name = '#hospital' + value.Record_id; 
//  var hospital_id = '"hospital' + value.Record_id +'"';
    var diagnosis = 'id="diagnosis' + value.Record_id + '"';
    var diagnosis_name = '#diagnosis' + value.Record_id;
//  var diagnosis_id = '"diagnosis' + value.Record_id + '"';
    var procedure = 'id="procedure' + value.Record_id + '"';
    var procedure_name = '#procedure' + value.Record_id;
//  var procedure_id = '"procedure' + value.Record_id + '"';
    var gender = 'id="gender' + value.Record_id + '"';
    var record_birthday = 'id="record_birthday' + value.Record_id + '"';
    var record_date = 'id="record_date' + value.Record_id + '"';
    var record_id = 'id="record_draft_id' + value.Record_id + '"';
    var record_id_name = '#record_draft_id' + value.Record_id;
    var result = 'id="result' + value.Record_id + '"';
    var special_notes = 'id="special_notes' + value.Record_id + '"';
    var weight = 'id="weight' + value.Record_id + '"';
    var delete_draft = 'id="delete_draft' + value.Record_id + '"';
    var delete_draft_name = '#delete_draft' + value.Record_id;

    info += '<tr>';

    info += '<td>' + 
        value.Record_id + '</td>';

    info += '<td>' + 
        value.Hospital_name + '</td>';

    info += '<td>' + 
        value.Start_datetime + '</td>';

    info += '<td>' + 
        value.Medicalemployee_firstname + '</td>';
        
    info += '<td>' + 
        value.Medicalemployee_lastname + '</td>';
    
    info += '<td>' + 
        value.Procedure_name + '</td>';
        
    info += '<td>' + 
    value.Diagnosis_name + '</td>';
    
    info += '<td>' + 
    value.Outcome + '</td>';

    info += `<td>
    <button type="button" class="btn btn-primary" data-toggle="modal" data-target="`+ modalcreate_name_edit +`">Edit</button>
    <input type="button" class="btn btn-danger" ` + delete_draft + ` value="Delete">
    </td>`;

    info += '</tr>'

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
                    <input type="text" class="form-control" `+ record_id +`name="record_draft_id" value="`+ value.Record_id +`" readonly>                
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
            <button type="submit" id="draft_submit_delete" class="btn btn-primary">Submit</button>
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

        $('#modal_editinfo').append(modal_editinfo);
        $(modalbody_name_edit).append(editinfo);

        $(document).ready(function() {
            $(delete_draft_name).click(function () {
                //console.log(value.Record_id);
                $.ajax(  
                    {
                        url:'/delete_record_draft',    
                        type:"POST",    
                        data: JSON.stringify({"Record_draft_id": value.Record_id}),
                        success:function(){  
                            //console.log("Successfully deleted!");
                            },
                        error: function(XMLHttpRequest, textStatus, errorThrown) { 
                            //console.log("Status: " + textStatus); 
                            //console.log("Error: " + errorThrown); 
                        }
                    });
            });
          });

 });
 $('#record_draft_datatable').append(info);
});



